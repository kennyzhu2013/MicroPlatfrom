package sync

import (
	"golang.org/x/net/context"
)

type syncKey struct{}

func FromContext(ctx context.Context) (Sync, bool) {
	c, ok := ctx.Value(syncKey{}).(Sync)
	return c, ok
}

func NewContext(ctx context.Context, c Sync) context.Context {
	return context.WithValue(ctx, syncKey{}, c)
}
