package log

import (
	"golang.org/x/net/context"
)

type Options struct {
	// the current log level
	Level Level
	// the output to write to
	Outputs []Output
	// include a set of fields
	Fields Fields

	// Alternative options
	Context context.Context
}

type OutputOptions struct {
	// filepath, url, etc
	Name string
}

// Log options

func WithLevel(l Level) Option {
	return func(o *Options) {
		o.Level = l
	}
}

func WithOutput(ot Output) Option {
	return func(o *Options) {
		o.Outputs = append(o.Outputs, ot)
	}
}

func WithFields(f Fields) Option {
	return func(o *Options) {
		o.Fields = f
	}
}

// Output options

func OutputName(name string) OutputOption {
	return func(o *OutputOptions) {
		o.Name = name
	}
}
