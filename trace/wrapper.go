package trace

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"time"

	"golang.org/x/net/context"
)

type clientWrapper struct {
	client.Client
	t Trace
	s *registry.Service
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	var span *Span
	var ok, okk bool
	var err error

	md, mk := metadata.FromContext(ctx)
	if !mk {
		md = make(metadata.Metadata)
	}

	// try pull span from context
	span, ok = SpanFromContext(ctx)
	if !ok {
		if !mk {
			// metadata is new, let's create a new span
			span = c.t.NewSpan(nil)
		} else {
			// ok we got some md

			// can we get the span from the header?
			span, okk = SpanFromHeader(md)
			if !okk {
				// no, ok create one!
				span = c.t.NewSpan(nil)
			}
		}
	}

	// got parent span from context or metadata
	if okk || ok {
		// setup the span with parent
		span = c.t.NewSpan(&Span{
			// same trace id
			TraceId: span.TraceId,
			// set parent id to parent span id
			ParentId: span.Id,
			// use previous debug
			Debug: span.Debug,
		})
	}

	// start the span
	span.Annotations = append(span.Annotations, &Annotation{
		Timestamp: time.Now(),
		Type:      AnnStart,
		Service:   c.s,
	})

	// and mark as debug? might want to do this based on a setting
	span.Debug = true
	// set uniq span name
	span.Name = req.Service() + "." + req.Method()
	// set source/dest
	span.Source = c.s
	span.Destination = &registry.Service{Name: req.Service()}

	// set context key
	newCtx := ContextWithSpan(ctx, span)
	// set metadata
	newCtx = metadata.NewContext(newCtx, HeaderWithSpan(md, span))

	// mark client request
	span.Annotations = append(span.Annotations, &Annotation{
		Timestamp: time.Now(),
		Type:      AnnClientRequest,
		Service:   c.s,
	})

	// defer the completion of the span
	defer func() {
		// mark client response
		span.Annotations = append(span.Annotations, &Annotation{
			Timestamp: time.Now(),
			Type:      AnnClientResponse,
			Service:   c.s,
		})

		// if we were the creator
		var debug map[string]string
		if err != nil {
			debug = map[string]string{"error": err.Error()}
		}
		// mark end of span
		span.Annotations = append(span.Annotations, &Annotation{
			Timestamp: time.Now(),
			Type:      AnnEnd,
			Service:   c.s,
			Debug:     debug,
		})

		span.Duration = time.Now().Sub(span.Timestamp)

		// flush the span to the collector on return
		c.t.Collect(span)
	}()

	// now just make a regular call down the stack
	err = c.Client.Call(newCtx, req, rsp, opts...)
	return err
}

func handlerWrapper(fn server.HandlerFunc, t Trace, s *registry.Service) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		// embed trace instance
		newCtx := NewContext(ctx, t)

		var span *Span
		var err error

		// get trace info from metadata
		md, ok := metadata.FromContext(ctx)
		if !ok {
			// this is a new span
			span = t.NewSpan(nil)
		} else {
			// can we gt the span from the header?
			span, ok = SpanFromHeader(md)
			if !ok {
				// no, ok create one
				span = t.NewSpan(nil)
			}
		}

		// mark client request
		span.Annotations = append(span.Annotations, &Annotation{
			Timestamp: time.Now(),
			Type:      AnnServerRequest,
			Service:   s,
		})

		// and mark as debug? might want to do this based on a setting
		span.Debug = true
		// set unique span name
		span.Name = req.Service() + "." + req.Method()
		// set source/dest
		span.Source = s
		span.Destination = &registry.Service{Name: req.Service()}

		// embed the span in the context
		newCtx = ContextWithSpan(newCtx, span)

		// defer the completion of the span
		defer func() {
			var debug map[string]string
			if err != nil {
				debug = map[string]string{"error": err.Error()}
			}
			// mark server response
			span.Annotations = append(span.Annotations, &Annotation{
				Timestamp: time.Now(),
				Type:      AnnServerResponse,
				Service:   s,
				Debug:     debug,
			})

			span.Duration = time.Now().Sub(span.Timestamp)

			// flush the span to the collector on return
			t.Collect(span)
		}()
		err = fn(newCtx, req, rsp)
		return err
	}
}
