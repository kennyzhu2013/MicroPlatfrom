package log

import (
	"golang.org/x/net/context"
)

type Options struct {
	// the current log level
	Level Level

	// the output to write to
	Op Output

	// include a set of fields
	Fields Fields
    // size of one log file

    FileSize int

	// Alternative options
	Context context.Context

	//
	OpOption OutputOptions
}

type OutputOptions struct {
	// file path, url, etc, Dir default is ""
	Name string
	Dir  string
}

// Log options

func WithLevel(l Level) Option {
	return func(o *Options) {
		o.Level = l
	}
}

func WithOutput(ot Output) Option {
	return func(o *Options) {
		o.Op = ot
	}
}

func WithFields(f Fields) Option {
	return func(o *Options) {
		o.Fields = f
	}
}

// Output options
func WithName(name string) Option {
	return func(o *Options) {
		o.OpOption.Name = name
	}
}

func WithDir(dir string) Option {
	return func(o *Options) {
		o.OpOption.Dir = dir
	}
}

// Output options
func OutputName(name string) OutputOption {
	return func(o *OutputOptions) {
		o.Name = name
	}
}

func OutputDir(dir string) OutputOption {
	return func(o *OutputOptions) {
		o.Dir = dir
	}
}
