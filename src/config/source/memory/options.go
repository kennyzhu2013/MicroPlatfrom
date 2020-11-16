package memory

import (
	"github.com/micro/go-os/config"

	"golang.org/x/net/context"
)

type dataKey struct{}

func Data(d []byte) config.SourceOption {
	return func(o *config.SourceOptions) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, dataKey{}, d)
	}
}
