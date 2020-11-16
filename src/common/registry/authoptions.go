package registry

import (
	"context"
)

type authKey struct{}

type authCreds struct {
	Username string
	Password string
}

// Auth allows you to specify username/password
func Auth(username, password string) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, authKey{}, &authCreds{Username: username, Password: password})
	}
}
