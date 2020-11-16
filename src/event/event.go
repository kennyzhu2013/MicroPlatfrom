// Package event is an interface for event sourcing.
package event

import (
	"golang.org/x/net/context"
)

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

type Handler func(*Record)

var (
	RecordTopic      = "micro.event.record"
	DefaultEventType = "event"
)

type Option func(o *Options)

func NewEvent(opts ...Option) Event {
	return newOS(opts...)
}
