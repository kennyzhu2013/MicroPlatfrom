// Package router is an interface for global service routing.
package router

import (
	"time"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
)

// The router is the client interface to a
// global service loadbalancer (GSLB).
// Metrics are batched and published to
// a router which has a view of the whole
// system.
type Router interface {
	// Provides the selector interface
	selector.Selector
	// Return stats maintained here
	Stats() ([]*Stats, error)
	// Record stats for a request - too many args ugh
	Record(r Request, node *registry.Node, d time.Duration, err error)
}

type Stats struct {
	Service   *registry.Service
	Client    *registry.Service
	Timestamp int64
	Duration  int64
	// TODO:
	// Selected
	// Endpoints
}

// Could be client or server request
type Request interface {
	Service() string
	Method() string
}

var (
	StatsTopic  = "micro.router.stats"
	LabelPrefix = "X-Micro-Label-"
)

func NewRouter(opts ...selector.Option) Router {
	return newOS(opts...)
}
