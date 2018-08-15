# Auth [![GoDoc](https://godoc.org/github.com/micro/go-os?status.svg)](https://godoc.org/github.com/micro/go-os/auth)

Provides a high level pluggable abstraction for authentication and authorisation.

## Interface

Simplify authentication with an interface that just returns true or 
false based on the current RPC context or session id. Optionally 
returns the session information for further examination.

Granular role based authorisation and control is needed at large scale 
for access management. Goes beyond just, does this person have an 
authenticated session. Should they be allowed to access the given 
resource.

Management of auth/roles should be offloaded to a service to minimise code changes 
in each individual service. Should ideally be embedded as middleware in requests handlers 
and initialised when registering a handler.

```go
// Auth handles client side validation of authentication
// The client does not actually handle authentication itself.
// This could be an oauth2 provider, openid, basic auth, etc.
type Auth interface {
	Authorized(ctx context.Context, req Request) (*Token, error)
	Introspect(ctx context.Context) (*Token, error)
	Revoke(t *Token) error
	Token() (*Token, error)
	String() string
}

func NewAuth(opts ...Option) Auth {
	return newOS(opts...)
}
```

## Supported Backends

- [Auth service](https://github.com/micro/auth-srv) (Oauth2)
