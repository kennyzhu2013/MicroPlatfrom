package log

import (
	"encoding/json"
	los "os"
	"reflect"
	"runtime"
	"strconv"

	"sync"
	"sort"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"strings"
)

type output struct {
	opts OutputOptions
	dirs string
	err error

	// json file encoder.
	f   *los.File
	mu   sync.Mutex

	startRoll sync.Once
	rollCh    chan bool
}

// logInfo is a convenience struct to return the filename and its embedded
// timestamp.
type logInfo struct {
	sequence int
	los.FileInfo
}


func (o *output) Send(e *Event) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.f == nil {
		return o.err
	}

	filename := o.f.Name()
	info, err := los.Stat(filename)
	if los.IsNotExist(err) {
		o.f,_ = los.OpenFile(filename, los.O_CREATE|los.O_APPEND|los.O_WRONLY, 0666)
	} else {
		// judge is full.
		if info.Size() >= FileSize {
			// rotate directly
			// o.rotate()
			if err = o.rollRunOnce();err != nil {
				_ = fmt.Errorf("rollRunOnce failed: %s", err)
			}
		}
	}

	return json.NewEncoder(o.f).Encode(e)
}

func (o *output) flush() error {
	if o.f == nil {
		return o.err
	}
	return o.f.Sync()
}

func (o *output) Close() error {
	o.mu.Lock()
	defer o.mu.Unlock()
	return o.close()
}

func (o *output) String() string {
	return o.opts.Name
}

func (o *output) rollRun() {
	for _ = range o.rollCh {
		// log?..
		_ = o.rollRunOnce()
	}
}

// close closes the file if it is open.
func (o *output) close() error {
	if o.f == nil {
		return nil
	}

	o.flush()
	err := o.f.Close()
	o.f = nil
	return err
}

// openNew opens a new log file for writing, moving any old log file out of the
// way.  This methods assumes the file has already been closed.
func (o *output) openNew() error {
	f, err := los.OpenFile(o.dirs + "/" +o.opts.Name, los.O_CREATE|los.O_APPEND|los.O_WRONLY, 0666)
	o.f = f
	fmt.Println()
	return err
}

// if file is full, move..
func (o *output)rotate()  error {
	// sync run.
	o.startRoll.Do(func() {
		o.rollCh = make(chan bool, 1)
		go o.rollRun()
	})
	select {
	case o.rollCh <- true:
	default:
	}

	return nil
}

func (o *output) rollRunOnce() error {
	files, err := o.oldLogFiles()
	if err != nil  || len(files) < 1 {
		return err
	}

	if err := o.close(); err != nil {
		return err
	}

	dir := o.dir()
	for _, f := range files {
		if f.sequence + 1 >= DefaultFileMaxNum {
			errRemove := los.Remove(filepath.Join(dir, f.Name()))
			if err == nil && errRemove != nil {
				err = errRemove
			}
			fmt.Printf("Remove file:%s\n", filepath.Join(dir, f.Name()))
			continue
		}

		errMove := los.Rename(filepath.Join(dir, f.Name()),
			filepath.Join(dir, o.nextFileName(f.sequence)))
		if err == nil && errMove != nil {
			err = errMove
		}
		fmt.Printf("Moved new file:source [%s], new [%s]\n", filepath.Join(dir, f.Name()), filepath.Join(dir, o.nextFileName(f.sequence)))
	}

	// after rotate, then create new.
	if err := o.openNew(); err != nil {
		return err
	}

	return err
}

func (o *output) nextFileName(i int) string {
	return o.opts.Name + "." + strconv.Itoa(i + 1)
}

// dir returns the directory for the current filename.
func (o *output) dir() string {
	if o.dirs == "" {
		o.dirs = filepath.Dir( o.f.Name() )
	}
	return o.dirs
}

// oldLogFiles returns the list of backup log files stored in the same
// directory as the current log file, sorted by ModTime
func (o *output) oldLogFiles() ([]logInfo, error) {
	files, err := ioutil.ReadDir(o.dir())
	if err != nil {
		return nil, fmt.Errorf("can't read log file directory: %s", err)
	}
	logFiles := []logInfo{}

	prefix := o.opts.Name

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if strings.HasPrefix(f.Name(), prefix) {
			filename := f.Name()
			var sequence int = 0
			if len(filename) > len(prefix) {
				ts := filename[len(prefix)+1 : ]
				if ts != "" {
					sequence,_ = strconv.Atoi(ts)
				}
			}

			logFiles = append(logFiles, logInfo{sequence, f})
			fmt.Printf("Found log file:name [%s], sequence [%d]\n",f.Name(), sequence)
			continue
		}

		// error parsing means that the suffix at the end was not generated
		// by lumberjack, and therefore it's not a backup file.
	}

	sort.Sort( byFormatSequence(logFiles) )
	return logFiles, nil
}

// byFormatTime sorts by newest time formatted in the name.
type byFormatSequence []logInfo

func (b byFormatSequence) Less(i, j int) bool {
	return b[i].sequence > b[j].sequence
}

func (b byFormatSequence) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byFormatSequence) Len() int {
	return len(b)
}

// 用 seps 进行分割, 根据协议栈信息查找...
func GetFunctionName(i interface{}, seps ...rune) string {
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()

	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})

	// fmt.Println(fields)

	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}


func NewOutput(opts ...OutputOption) Output {
	var options OutputOptions
	for _, o := range opts {
		o(&options)
	}

	if len(options.Name) == 0 {
		options.Name = DefaultOutputName
	}

	// move files..
	var logDir string
	if options.Dir != "" {
		logDir = options.Dir + "/" + options.Name
	}else {
		logDir = options.Name
	}
	f, err := los.OpenFile(logDir, los.O_CREATE|los.O_APPEND|los.O_WRONLY, 0666)

	return &output{
		opts: options,
		err:  err,
		f:    f,
		dirs: filepath.Dir( logDir ),
	}
}

