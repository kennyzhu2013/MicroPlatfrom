package auth

import (
	"golang.org/x/net/context"
)

type authKey struct{}
type tokenKey struct{}

func FromContext(ctx context.Context) (Auth, bool) {
	c, ok := ctx.Value(authKey{}).(Auth)
	return c, ok
}

func NewContext(ctx context.Context, c Auth) context.Context {
	return context.WithValue(ctx, authKey{}, c)
}

func TokenFromContext(ctx context.Context) (*Token, bool) {
	t, ok := ctx.Value(tokenKey{}).(*Token)
	return t, ok
}

func ContextWithToken(ctx context.Context, t *Token) context.Context {
	return context.WithValue(ctx, tokenKey{}, t)
}
