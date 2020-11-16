/*
@Time : 2019/4/19 16:24 
@Author : kenny zhu
@File : options.go
@Software: GoLand
@Others:
*/
package web

import (
	"context"
	"registry"
	"selector"
)

type Options struct {
	Registry registry.Registry
	Selector selector.Strategy
	Destination   string
	RouteTag	  string

	// Other options for implementations can be stored in a context. like selector_ttl.
	Context context.Context
}

type Option func(*Options)
func WithRegistry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

func WithSelector(s selector.Strategy) Option {
	return func(o *Options) {
		o.Selector = s
	}
}

func WithDestination(d string) Option {
	return func(o *Options) {
		o.Destination = d
	}
}

func WithRouteTag(d string) Option {
	return func(o *Options) {
		o.RouteTag = d
	}
}