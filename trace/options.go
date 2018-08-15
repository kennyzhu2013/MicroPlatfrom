package trace

import (
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
)

type Options struct {
	BatchSize      int
	BatchInterval  time.Duration
	CollectTimeout time.Duration
	Collectors     []string
	Topic          string

	Client  client.Client
	Service *registry.Service
}

func BatchSize(b int) Option {
	return func(o *Options) {
		o.BatchSize = b
	}
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

func Topic(t string) Option {
	return func(o *Options) {
		o.Topic = t
	}
}

func Service(s *registry.Service) Option {
	return func(o *Options) {
		o.Service = s
	}
}

func Client(c client.Client) Option {
	return func(o *Options) {
		o.Client = c
	}
}

func CollectTimeout(t time.Duration) Option {
	return func(o *Options) {
		o.CollectTimeout = t
	}
}
