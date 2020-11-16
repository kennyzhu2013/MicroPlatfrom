package discovery

import (
	"golang.org/x/net/context"
)

type discoveryKey struct{}

func FromContext(ctx context.Context) (Discovery, bool) {
	c, ok := ctx.Value(discoveryKey{}).(Discovery)
	return c, ok
}

func NewContext(ctx context.Context, c Discovery) context.Context {
	return context.WithValue(ctx, discoveryKey{}, c)
}
