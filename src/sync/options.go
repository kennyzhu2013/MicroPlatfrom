package sync

import (
	"time"

	"github.com/micro/go-micro/registry"
)

type Options struct {
	Namespace string
	Service   *registry.Service
	Nodes     []string
}

type LockOptions struct {
	TTL  time.Duration
	Wait time.Duration
}

type LeaderOptions struct{}

func Namespace(n string) Option {
	return func(o *Options) {
		o.Namespace = n
	}
}

func Nodes(nodes ...string) Option {
	return func(o *Options) {
		o.Nodes = append(o.Nodes, nodes...)
	}
}

func Service(s *registry.Service) Option {
	return func(o *Options) {
		o.Service = s
	}
}

func LockTTL(t time.Duration) LockOption {
	return func(o *LockOptions) {
		o.TTL = t
	}
}

func LockWait(t time.Duration) LockOption {
	return func(o *LockOptions) {
		o.Wait = t
	}
}
