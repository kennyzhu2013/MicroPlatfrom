package router

import (
	"sync"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
)

// local node cache
type cache struct {
	// what we got back
	services []*registry.Service

	// valid until
	// -1 indicates never expires
	// 0 indicates invalid list
	// > 0 indicates expires sometime in future
	// expected unix timestamp
	expires int64

	// list of nodes and pointer to
	// current position
	sync.Mutex
	pointer int
	nodes   map[string]bool
}

func newCache(services []*registry.Service, expires int64) *cache {
	cache := &cache{
		services: services,
		expires:  expires,
		nodes:    make(map[string]bool),
	}

	for _, service := range services {
		for _, node := range service.Nodes {
			cache.nodes[node.Id] = true
		}
	}

	return cache
}

func (c *cache) Filter(filters []selector.Filter) (selector.Next, error) {
	services := c.services

	for _, filter := range filters {
		services = filter(services)
	}

	if len(services) == 0 {
		return nil, selector.ErrNotFound
	}

	var nodes []*registry.Node

	for _, service := range services {
		for _, node := range service.Nodes {
			nodes = append(nodes, node)
		}
	}

	if len(nodes) == 0 {
		return nil, selector.ErrNotFound
	}

	// TODO: picking algorithm based on what the router returns
	// roundrobin
	// random
	// leastconn

	c.Lock()
	p := c.pointer
	if c.pointer >= len(c.nodes) {
		c.pointer = 0
	} else {
		c.pointer++
	}
	c.Unlock()

	return func() (*registry.Node, error) {
		if p >= len(nodes) {
			p = 0
		}

		node := nodes[p]
		p++

		return node, nil
	}, nil
}
