// Package auth is an interface for authentication and authorization.
package auth

import (
	"errors"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"

	"golang.org/x/net/context"
)

//https://github.com/beego/social-auth
// Auth handles client side validation of authentication
// The client does not actually handle authentication itself.
// This could be an oauth2 provider, openid, basic auth, etc.
type Auth interface {
	// Check if authorised
	Authorized(ctx context.Context, req Request) (*Token, error)
	// Retrieve token from context
	Introspect(ctx context.Context) (*Token, error)
	// Revoke a token
	Revoke(t *Token) error
	// Retrieve client token
	Token() (*Token, error)
	String() string
}

////https://github.com/beego/social-auth;TODO..

type Option func(*Options)

// Client or server request
type Request interface {
	Service() string
	Method() string
}

// Basically identical to oauth token
type Token struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	ExpiresAt    time.Time
	Scopes       []string
	Metadata     map[string]string
}

var (
	ErrInvalidToken = errors.New("invalid token")
)

func ClientWrapper(a Auth) client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c, a}
	}
}

func HandlerWrapper(a Auth) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return handlerWrapper(h, a)
	}
}

func NewAuth(opts ...Option) Auth {
	return newOS(opts...)
}
