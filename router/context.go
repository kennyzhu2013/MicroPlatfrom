package router

import (
	"golang.org/x/net/context"
)

type routerKey struct{}

func FromContext(ctx context.Context) (Router, bool) {
	c, ok := ctx.Value(routerKey{}).(Router)
	return c, ok
}

func NewContext(ctx context.Context, c Router) context.Context {
	return context.WithValue(ctx, routerKey{}, c)
}
