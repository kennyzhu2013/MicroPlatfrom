package monitor

import (
	"golang.org/x/net/context"
)

type monitorKey struct{}

func FromContext(ctx context.Context) (Monitor, bool) {
	c, ok := ctx.Value(monitorKey{}).(Monitor)
	return c, ok
}

func NewContext(ctx context.Context, c Monitor) context.Context {
	return context.WithValue(ctx, monitorKey{}, c)
}
