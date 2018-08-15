// Package trace is an interface for tracing.
package trace

import (
	"sync"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
)

const (
	AnnUnknown            AnnotationType = 0
	AnnStart              AnnotationType = 1
	AnnEnd                AnnotationType = 2
	AnnTimeout            AnnotationType = 3
	AnnClientRequest      AnnotationType = 4
	AnnClientResponse     AnnotationType = 5
	AnnClientPublication  AnnotationType = 6
	AnnServerRequest      AnnotationType = 7
	AnnServerResponse     AnnotationType = 8
	AnnServerSubscription AnnotationType = 9
)

type AnnotationType int32

type Trace interface {
	Close() error
	// New span with certain fields preset.
	// Provide parent span if you have it.
	NewSpan(*Span) *Span
	// Collect spans
	Collect(*Span) error
	// Name
	String() string
}

type Span struct {
	Name      string        // Topic / RPC Method
	Id        string        // id of this span
	TraceId   string        // The root trace id
	ParentId  string        // Parent span id
	Timestamp time.Time     // Microseconds from epoch. When span started.
	Duration  time.Duration // Microseconds. Duration of the span.
	Debug     bool          // Should persist no matter what.

	Source      *registry.Service // Originating service
	Destination *registry.Service // Destination service

	sync.Mutex
	Annotations []*Annotation
}

type Annotation struct {
	Timestamp time.Time // Microseconds from epoch
	Type      AnnotationType
	Key       string
	Value     []byte
	Debug     map[string]string
	Service   *registry.Service // Annotator
}

type Option func(o *Options)

var (
	DefaultBatchSize     = 100
	DefaultBatchInterval = time.Second * 5

	TraceTopic   = "micro.trace.span"
	TraceHeader  = "X-Micro-Trace-Id"
	SpanHeader   = "X-Micro-Span-Id"
	ParentHeader = "X-Micro-Parent-Id"
	DebugHeader  = "X-Micro-Trace-Debug"
)

func ClientWrapper(t Trace, s *registry.Service) client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c, t, s}
	}
}

func HandlerWrapper(t Trace, s *registry.Service) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return handlerWrapper(h, t, s)
	}
}

func NewTrace(opts ...Option) Trace {
	return newOS(opts...)
}
