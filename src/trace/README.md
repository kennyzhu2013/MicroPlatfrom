# Trace [![GoDoc](https://godoc.org/github.com/micro/go-os?status.svg)](https://godoc.org/github.com/micro/go-os/trace)

Provides a pluggable distributed tracing interface

```go
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

func NewTrace(opts ...Option) Trace {
	return newOS(opts...)
}
```

## Supported

- [Trace service](https://github.com/micro/trace-srv)
- [Zipkin](https://github.com/micro/go-plugins/tree/master/trace/zipkin)
- [Open Tracing](https://github.com/micro/go-plugins/tree/master/wrapper/trace/opentracing)

## Usage

You can either manually use the Trace interface to create and collect spans - look at the [wrappers](https://github.com/micro/go-os/blob/master/trace/wrapper.go) 
for an example - or use the client and server wrappers which will be called on every request made or received.

Also check out [go-os/examples/trace](https://github.com/micro/go-os/tree/master/examples/trace).

```go
import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-os/trace"
)

func main() {
	t := trace.NewTrace()

	srv := &registry.Service{Name: "go.micro.srv.example"}

	service := micro.NewService(
		micro.Name("go.micro.srv.example"),
		micro.WrapClient(trace.ClientWrapper(t, srv)),
		micro.WrapHandler(trace.HandlerWrapper(t, srv)),
	)
}
```

