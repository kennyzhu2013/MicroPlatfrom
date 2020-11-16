package router

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/selector"
	"github.com/micro/go-micro/server"
)

func Client(c client.Client) selector.Option {
	return func(o *selector.Options) {
		o.Context = client.NewContext(o.Context, c)
	}
}

func Server(s server.Server) selector.Option {
	return func(o *selector.Options) {
		o.Context = server.NewContext(o.Context, s)
	}
}
