package log

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

type os struct {
	// fn logFunc
	opts Options
}

/*
type logger struct {
	f  Fields
	fn logFunc
}*/

type logFunc func(l Level, f Fields, m string) error

func newOS(opts ...Option) Log {
	options := Options{
		Level:   DefaultLevel,
		Fields:  make(Fields),
		Context: context.TODO(),
	}

	for _, o := range opts {
		o(&options)
	}

	if options.Op == nil {
		options.Op = NewOutput( OutputDir(options.OpOption.Dir), OutputName(options.OpOption.Name) )
	}

	// https://blog.csdn.net/linuxandroidwince/article/details/79113398...
	o := &os{
		options,
	}

	return o
}

// use logger fields now.
func (o *os) log(l Level, m string) error {
	// discard if we're not at the right level
	if l < o.opts.Level {
		return nil
	}

	e := &Event{
		Timestamp: time.Now().Format("2006-1-2 15:04:05"),
		Level:     l,
		Fields:    o.opts.Fields,
		Message:   m,
	}

	// write directly.
	if err := o.opts.Op.Send(e); err != nil {
		return err
	}

	return nil
}

func (o *os) Init(opts ...Option) error {
	for _, opt := range opts {
		opt(&o.opts)
	}
	return nil
}

func (o *os) Options() Options {
	return o.opts
}

func (o *os) String() string {
	return "os"
}

func (o *os) Debug(args ...interface{}) {
	o.log(DebugLevel, fmt.Sprint(args...))
}

func (o *os) Info(args ...interface{}) {
	o.log(InfoLevel, fmt.Sprint(args...))
}

func (o *os) Warn(args ...interface{}) {
	o.log(WarnLevel, fmt.Sprint(args...))
}
func (o *os) Error(args ...interface{}) {
	o.log(ErrorLevel, fmt.Sprint(args...))
}

func (o *os) Fatal(args ...interface{}) {
	o.log(FatalLevel, fmt.Sprint(args...))
}

func (o *os) Debugf(format string, args ...interface{}) {
	o.log(DebugLevel, fmt.Sprintf(format, args...))
}

func (o *os) Infof(format string, args ...interface{}) {
	o.log(InfoLevel, fmt.Sprintf(format, args...))
}

func (o *os) Warnf(format string, args ...interface{}) {
	o.log(WarnLevel, fmt.Sprintf(format, args...))
}

func (o *os) Errorf(format string, args ...interface{}) {
	o.log(ErrorLevel, fmt.Sprintf(format, args...))
}

func (o *os) Fatalf(format string, args ...interface{}) {
	o.log(FatalLevel, fmt.Sprintf(format, args...))
}

func (o *os) Log(level Level, args ...interface{}) {
	o.log(level, fmt.Sprint(args...))
}

func (o *os) Logf(level Level, format string, args ...interface{}) {
	o.log(level, fmt.Sprintf(format, args...))
}

func (o *os) WithFields(f Fields) Logger {
	options := o.opts

	for k, v := range o.opts.Fields {
		options.Fields[k] = v
	}

	for k, v := range f {
		options.Fields[k] = v
	}

	return &os{
		opts:  options,
	}
}
