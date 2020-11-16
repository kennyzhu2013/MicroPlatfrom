# Event [![GoDoc](https://godoc.org/github.com/micro/go-os?status.svg)](https://godoc.org/github.com/micro/go-os/event)

A pluggable interface for platform event publishing and subscription.

## Interface

It's important to be able to track events at the platform level to know what's changing where. With 
hundreds of services it can be difficult to track or find a way to enforce this across services. 
By tracking platform events we can essentially build a platform event correlation system.

```go
type Event interface {
	// publish an event record
	Publish(context.Context, *Record) error
	// subscribe to an event types
	Subscribe(context.Context, Handler, ...string) error
	// Name
	String() string
}

type Record struct {
	Id        string
	Type      string
	Origin    string
	Timestamp int64
	RootId    string
	Metadata  map[string]string
	Data      string
}

func NewEvent(opts ...Option) Event {
	return newOS(opts...)
}
```

## Supported Backends
- [Event service](https://github.com/micro/event-srv)
