// Package discovery is an interface for scalable service discovery.
package discovery

import (
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry"
)

const (
	HeartbeatTopic = "micro.discovery.heartbeat"
)

// Discovery builds on the registry for heartbeating and client side caching
type Discovery interface {
	Close() error
	registry.Registry
}

func init() {
	cmd.DefaultRegistries["os"] = NewRegistry
}

func NewDiscovery(opts ...registry.Option) Discovery {
	return newOS(opts...)
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	return newOS(opts...)
}
