/*
@Time : 2019/4/23 15:37 
@Author : kenny zhu
@File : options.go
@Software: GoLand
@Others:
*/
package monitor


import (
"time"

"github.com/micro/go-micro/registry"
"context"
)

type optionsKey struct{}

// Options holds discovery options stored in registry Options context.
type Options struct {
	Interval time.Duration
}

func getOptions(ctx context.Context) *Options {
	if opts, ok := ctx.Value(optionsKey{}).(*Options); ok {
		return opts
	}
	return &Options{}
}

func setOptions(ctx context.Context, opts *Options) context.Context {
	return context.WithValue(ctx, optionsKey{}, opts)
}

// Interval on which to publish heartbeats
func WithInterval(i time.Duration) registry.Option {
	return func(o *registry.Options) {
		opts := getOptions(o.Context)
		opts.Interval = i
		o.Context = setOptions(o.Context, opts)
	}
}

