package discovery

import (
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"golang.org/x/net/context"
)

type optionsKey struct{}

// Options holds discovery options stored in registry Options context
type Options struct {
	Client   client.Client
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

// Client used to call the discovery service
func Client(c client.Client) registry.Option {
	return func(o *registry.Options) {
		opts := getOptions(o.Context)
		opts.Client = c
		o.Context = setOptions(o.Context, opts)
	}
}

// Interval on which to publish heartbeats
func Interval(i time.Duration) registry.Option {
	return func(o *registry.Options) {
		opts := getOptions(o.Context)
		opts.Interval = i
		o.Context = setOptions(o.Context, opts)
	}
}
