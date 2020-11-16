package metrics

import (
	"time"
)

type Options struct {
	Namespace     string
	Fields        Fields
	Collectors    []string
	BatchInterval time.Duration
}

func BatchInterval(b time.Duration) Option {
	return func(o *Options) {
		o.BatchInterval = b
	}
}

func Collectors(c ...string) Option {
	return func(o *Options) {
		o.Collectors = c
	}
}

func Namespace(n string) Option {
	return func(o *Options) {
		o.Namespace = n
	}
}

func WithFields(f Fields) Option {
	return func(o *Options) {
		o.Fields = f
	}
}
