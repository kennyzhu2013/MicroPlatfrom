package event

import (
	"github.com/micro/go-micro/client"

	"golang.org/x/net/context"
)

type Options struct {
	Client client.Client

	// Alternate options
	Context context.Context
}
