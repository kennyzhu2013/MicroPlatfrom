package consul

import (
	"crypto/md5"
	"fmt"
	"net"

	"github.com/hashicorp/consul/api"
	"github.com/micro/go-os/config"
)

// Currently a single consul reader
type consul struct {
	addr   string
	opts   config.SourceOptions
	client *api.Client
}

var (
	DefaultPath = "/micro/config"
)

func (c *consul) Read() (*config.ChangeSet, error) {
	kv, _, err := c.client.KV().Get(c.opts.Name, nil)
	if err != nil {
		return nil, err
	}

	if kv == nil {
		return nil, fmt.Errorf("source not found: %s", c.opts.Name)
	}

	// hash the consul
	h := md5.New()
	h.Write(kv.Value)
	checksum := fmt.Sprintf("%x", h.Sum(nil))

	return &config.ChangeSet{
		Source:   c.String(),
		Data:     kv.Value,
		Checksum: checksum,
	}, nil
}

func (c *consul) String() string {
	return "consul"
}

func (c *consul) Watch() (config.SourceWatcher, error) {
	w, err := newWatcher(c.opts.Name, c.addr, c.String())
	if err != nil {
		return nil, err
	}
	return w, nil
}

func NewSource(opts ...config.SourceOption) config.Source {
	options := config.SourceOptions{
		Name: DefaultPath,
	}

	for _, o := range opts {
		o(&options)
	}

	// use default config
	config := api.DefaultConfig()

	// check if there are any addrs
	if len(options.Hosts) > 0 {
		addr, port, err := net.SplitHostPort(options.Hosts[0])
		if ae, ok := err.(*net.AddrError); ok && ae.Err == "missing port in address" {
			port = "8500"
			addr = options.Hosts[0]
			config.Address = fmt.Sprintf("%s:%s", addr, port)
		} else if err == nil {
			config.Address = fmt.Sprintf("%s:%s", addr, port)
		}
	}

	// create the client
	client, _ := api.NewClient(config)

	return &consul{
		addr:   config.Address,
		opts:   options,
		client: client,
	}
}
