# Router [![GoDoc](https://godoc.org/github.com/micro/go-os?status.svg)](https://godoc.org/github.com/micro/go-os/router)

The router is a client library for global service load balancing. Go-micro uses client side load balancing 
with the `Selector` interface but most implementation only provide a single view point of the environment, 
from the service itself. The router library talks to a backend service which aggregates metrics from all 
services and relays back routing information.

<p align="center">
  <img src="https://github.com/micro/go-os/blob/master/doc/router.png" />
</p>

## Interface

```go
// The router is the client interface for 
// global service load balancing (GSLB).
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

func NewRouter(opts ...selector.Option) Router {
	return newOS(opts...)
}
```

## Supported Backends

- [Router service](https://github.com/micro/router-srv)
