package metrics

import (
	"golang.org/x/net/context"
)

type metricsKey struct{}

func FromContext(ctx context.Context) (Metrics, bool) {
	c, ok := ctx.Value(metricsKey{}).(Metrics)
	return c, ok
}

func NewContext(ctx context.Context, c Metrics) context.Context {
	return context.WithValue(ctx, metricsKey{}, c)
}
