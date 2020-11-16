package kv

import (
	"golang.org/x/net/context"
)

type kvKey struct{}

func FromContext(ctx context.Context) (KV, bool) {
	c, ok := ctx.Value(kvKey{}).(KV)
	return c, ok
}

func NewContext(ctx context.Context, c KV) context.Context {
	return context.WithValue(ctx, kvKey{}, c)
}
