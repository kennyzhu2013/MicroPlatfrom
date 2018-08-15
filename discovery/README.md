# Discovery [![GoDoc](https://godoc.org/github.com/micro/go-os?status.svg)](https://godoc.org/github.com/micro/go-os/discovery)

Discovery builds on the go-micro.Registry and is backed by the Micro OS discovery service.

The go-micro registry provides a simple Registry abstraction for various service discovery systems. 
Heartbeating is also done through a simple form of re-registration. Because of this we end up 
with a system that has limited scaling potential and does not provide much information about 
service health.

Discovery provides heartbeating as events published via the Broker. This means anyone can subscribe 
to the heartbeats to determine "liveness" rather than querying the Registry. Discovery also 
includes an in-memory cache of the Registry using the Watcher. If the Registry fails for any 
reason Discovery continues to function.

## Interface

```go
// Discovery builds on the registry as a mechanism
// for finding services. It includes heartbeating
// to notify of liveness and caching of the registry.
type Discovery interface {
	// implements the registry interface
	registry.Registry
	Close() error
}

func NewDiscovery(opts ...registry.Option) Discovery {
	return newOS(opts...)
}
```

## Supported Backends

- [Discovery Service](https://github.com/micro/discovery-srv)

## Usage

### With Flag

```go
import _ "github.com/micro/go-os/discovery"
```

```shell
go run main.go --registry=os --registry_address=addr1:port,addr2:port,addr3:port
```

### Direct Use

```go
import (
	"github.com/micro/go-micro"
	"github.com/micro/go-os/discovery"
)

func main() {
	r := discovery.NewDiscovery(
		discovery.Addrs("addr1:port", "addr2:port", "addr3:port"),
	)

	service := micro.NewService(
		micro.Name("my.service"),
		micro.Registry(r),
	)
}
```
