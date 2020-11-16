package etcd

import (
	"crypto/md5"
	"fmt"
	"log"

	"github.com/coreos/etcd/client"
	"github.com/micro/go-os/config"

	"golang.org/x/net/context"
)

type etcd struct {
	addrs  []string
	opts   config.SourceOptions
	client client.Client
}

var (
	DefaultPath = "/micro/config"
)

func (e *etcd) Read() (*config.ChangeSet, error) {
	kv := client.NewKeysAPI(e.client)

	rsp, err := kv.Get(context.Background(), e.opts.Name, nil)
	if err != nil {
		return nil, err
	}

	// hash the etcd
	h := md5.New()
	h.Write([]byte(rsp.Node.Value))
	checksum := fmt.Sprintf("%x", h.Sum(nil))

	return &config.ChangeSet{
		Source:   e.String(),
		Data:     []byte(rsp.Node.Value),
		Checksum: checksum,
	}, nil
}

func (e *etcd) String() string {
	return "etcd"
}

func (e *etcd) Watch() (config.SourceWatcher, error) {
	w, err := newWatcher(e.opts.Name, e.addrs, e.String())
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

	var cAddrs []string

	for _, addr := range options.Hosts {
		if len(addr) > 0 {
			cAddrs = append(cAddrs, addr)
		}
	}

	if len(cAddrs) == 0 {
		cAddrs = []string{"http://127.0.0.1:2379"}
	}

	c, err := client.New(client.Config{
		Endpoints: cAddrs,
	})

	if err != nil {
		log.Fatal(err)
	}

	return &etcd{
		addrs:  cAddrs,
		opts:   options,
		client: c,
	}
}
