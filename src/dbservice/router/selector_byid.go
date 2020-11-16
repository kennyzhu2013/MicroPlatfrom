/*
@Time : 2018/8/24 14:22 
@Author : kenny zhu
@File : selector_byid.go
@Software: GoLand
@Others: Implements multi-database for route services...
*/
package router

import (
	"github.com/micro/go-micro/selector"
	"math/rand"
	"time"
	"github.com/micro/go-micro/registry"
	"sync"
)

// Built in random hashed node selector
// One datasource must be named...
type idSelector struct {
	opts selector.Options
}

var (
	// All server the same tag in the same set..
	datacenter = "dbrouter"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func (n *idSelector) Init(opts ...selector.Option) error {
	for _, o := range opts {
		o(&n.opts)
	}
	return nil
}

func (n *idSelector) Options() selector.Options {
	return n.opts
}

// Select directly not use other Strategy
func (n *idSelector) Select(service string, opts ...selector.SelectOption) (selector.Next, error) {
	// Default is ConsulRegistry.
	services, err := n.opts.Registry.GetService(service)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, selector.ErrNotFound
	}

	var nodes []*registry.Node
	// Filter the nodes for datacenter marked by the server..
	for _, service := range services {
		for _, node := range service.Nodes {
			if node.Metadata["datacenter"] == datacenter {
				nodes = append(nodes, node)
			}
		}
	}

	if len(nodes) == 0 {
		return nil, selector.ErrNotFound
	}

	var i int
	var mtx sync.Mutex

	// Round bin..
	return func() (*registry.Node, error) {
		mtx.Lock()
		defer mtx.Unlock()
		i++
		return nodes[i%len(nodes)], nil
	}, nil
}

func (n *idSelector) Mark(service string, node *registry.Node, err error) {
	return
}

func (n *idSelector) Reset(service string) {
	return
}

func (n *idSelector) Close() error {
	return nil
}

func (n *idSelector) String() string {
	return datacenter
}

// Return a new first node selector
func IDSelector(opts ...selector.Option) selector.Selector {
	var sopts selector.Options
	for _, opt := range opts {
		opt(&sopts)
	}
	if sopts.Registry == nil {
		sopts.Registry = registry.DefaultRegistry
	}
	return &idSelector{sopts}
}