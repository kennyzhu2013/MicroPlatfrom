package auth

import (
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

type Options struct {
	Id     string
	Secret string
	Client client.Client
	// Used for alternative options
	Context context.Context
}

func Client(c client.Client) Option {
	return func(o *Options) {
		o.Client = c
	}
}

func Id(id string) Option {
	return func(o *Options) {
		o.Id = id
	}
}

func Secret(s string) Option {
	return func(o *Options) {
		o.Secret = s
	}
}
