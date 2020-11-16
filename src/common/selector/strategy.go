/*
@Time : 2019/6/19 14:15
@Author : kenny zhu
@File : strategy
@Software: GoLand
@Others:
*/
package selector

import (
	"math/rand"
	"registry"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random is a random strategy algorithm for node selection
func Random(services []*registry.Service) Next {
	var nodes []*registry.Node

	for _, service := range services {
		nodes = append(nodes, service.Nodes...)
	}

	return func() (*registry.Node, error) {
		if len(nodes) == 0 {
			return nil, ErrNoneAvailable
		}

		i := rand.Int() % len(nodes)
		return nodes[i], nil
	}
}


// RoundRobin is a round robin strategy algorithm for node selection,
// WARN : Here not supports go-routines.
func RoundRobin(services []*registry.Service) Next {
	var nodes []*registry.Node

	for _, service := range services {
		nodes = append(nodes, service.Nodes...)
	}

	var i = rand.Int()
	var mtx sync.Mutex

	return func() (*registry.Node, error) {
		if len(nodes) == 0 {
			return nil, ErrNoneAvailable
		}

		mtx.Lock()
		node := nodes[i%len(nodes)]
		i++
		mtx.Unlock()

		return node, nil
	}
}