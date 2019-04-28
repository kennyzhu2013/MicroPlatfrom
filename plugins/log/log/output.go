package log

import (
	"encoding/json"
	los "os"
	"strconv"
	"io"
)

type output struct {
	opts OutputOptions

	err error
	f   *los.File
}

func (o *output) Send(e *Event) error {
	if o.f == nil {
		return o.err
	}
	return json.NewEncoder(o.f).Encode(e)
}

func (o *output) Flush() error {
	if o.f == nil {
		return o.err
	}
	return o.f.Sync()
}

func (o *output) Close() error {
	if o.f == nil {
		return o.err
	}
	return o.f.Close()
}

func (o *output) String() string {
	return o.opts.Name
}

// sjlin 新增判断文件大小方法
func (o *output) IsFull() bool {
	outputSize, _ := o.f.Seek(io.SeekStart, io.SeekEnd)
	if outputSize > FileSize {
		return true
	}
	return false
}

func NewOutput(opts ...OutputOption) Output {
	var options OutputOptions
	for _, o := range opts {
		o(&options)
	}

	if len(options.Name) == 0 {
		options.Name = DefaultOutputName
	}
    FileNumber = FileNumber + 1   // 每次创建文件，序号 + 1
	f, err := los.OpenFile(options.Name + strconv.Itoa(FileNumber), los.O_CREATE|los.O_APPEND|los.O_WRONLY, 0666)

	return &output{
		opts: options,
		err:  err,
		f:    f,
	}
}
