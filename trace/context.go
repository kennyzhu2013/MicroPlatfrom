package trace

import (
	"golang.org/x/net/context"
)

type traceKey struct{}
type spanKey struct{}

func FromContext(ctx context.Context) (Trace, bool) {
	c, ok := ctx.Value(traceKey{}).(Trace)
	return c, ok
}

func NewContext(ctx context.Context, c Trace) context.Context {
	return context.WithValue(ctx, traceKey{}, c)
}

func SpanFromContext(ctx context.Context) (*Span, bool) {
	s, ok := ctx.Value(spanKey{}).(*Span)
	return s, ok
}

func ContextWithSpan(ctx context.Context, s *Span) context.Context {
	return context.WithValue(ctx, spanKey{}, s)
}
