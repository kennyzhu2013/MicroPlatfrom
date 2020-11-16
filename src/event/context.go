package event

import (
	"golang.org/x/net/context"
)

type eventKey struct{}

func FromContext(ctx context.Context) (Event, bool) {
	c, ok := ctx.Value(eventKey{}).(Event)
	return c, ok
}

func NewContext(ctx context.Context, c Event) context.Context {
	return context.WithValue(ctx, eventKey{}, c)
}
