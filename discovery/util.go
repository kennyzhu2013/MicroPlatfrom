package discovery

import (
	"sync"

	"github.com/micro/go-micro/registry"
	proto "github.com/micro/go-os/discovery/proto"
)

func rr(addrs []string) func() []string {
	var i int
	var mtx sync.Mutex
	length := len(addrs)

	return func() []string {
		if length == 0 {
			return addrs
		}

		mtx.Lock()
		j := i
		i++
		mtx.Unlock()

		var newAddrs []string

		for k := 0; k < length; k++ {
			newAddrs = append(newAddrs, addrs[j%length])
			j++
		}

		return newAddrs
	}
}

func values(v []*registry.Value) []*proto.Value {
	if len(v) == 0 {
		return []*proto.Value{}
	}

	var vs []*proto.Value
	for _, vi := range v {
		vs = append(vs, &proto.Value{
			Name:   vi.Name,
			Type:   vi.Type,
			Values: values(vi.Values),
		})
	}
	return vs
}

func toValues(v []*proto.Value) []*registry.Value {
	if len(v) == 0 {
		return []*registry.Value{}
	}

	var vs []*registry.Value
	for _, vi := range v {
		vs = append(vs, &registry.Value{
			Name:   vi.Name,
			Type:   vi.Type,
			Values: toValues(vi.Values),
		})
	}
	return vs
}

func toProto(s *registry.Service) *proto.Service {
	var endpoints []*proto.Endpoint
	for _, ep := range s.Endpoints {
		var request, response *proto.Value

		if ep.Request != nil {
			request = &proto.Value{
				Name:   ep.Request.Name,
				Type:   ep.Request.Type,
				Values: values(ep.Request.Values),
			}
		}

		if ep.Response != nil {
			response = &proto.Value{
				Name:   ep.Response.Name,
				Type:   ep.Response.Type,
				Values: values(ep.Response.Values),
			}
		}

		endpoints = append(endpoints, &proto.Endpoint{
			Name:     ep.Name,
			Request:  request,
			Response: response,
			Metadata: ep.Metadata,
		})
	}

	var nodes []*proto.Node

	for _, node := range s.Nodes {
		nodes = append(nodes, &proto.Node{
			Id:       node.Id,
			Address:  node.Address,
			Port:     int64(node.Port),
			Metadata: node.Metadata,
		})
	}

	return &proto.Service{
		Name:      s.Name,
		Version:   s.Version,
		Metadata:  s.Metadata,
		Endpoints: endpoints,
		Nodes:     nodes,
	}
}

func toService(s *proto.Service) *registry.Service {
	var endpoints []*registry.Endpoint
	for _, ep := range s.Endpoints {
		var request, response *registry.Value

		if ep.Request != nil {
			request = &registry.Value{
				Name:   ep.Request.Name,
				Type:   ep.Request.Type,
				Values: toValues(ep.Request.Values),
			}
		}

		if ep.Response != nil {
			response = &registry.Value{
				Name:   ep.Response.Name,
				Type:   ep.Response.Type,
				Values: toValues(ep.Response.Values),
			}
		}

		endpoints = append(endpoints, &registry.Endpoint{
			Name:     ep.Name,
			Request:  request,
			Response: response,
			Metadata: ep.Metadata,
		})
	}

	var nodes []*registry.Node
	for _, node := range s.Nodes {
		nodes = append(nodes, &registry.Node{
			Id:       node.Id,
			Address:  node.Address,
			Port:     int(node.Port),
			Metadata: node.Metadata,
		})
	}

	return &registry.Service{
		Name:      s.Name,
		Version:   s.Version,
		Metadata:  s.Metadata,
		Endpoints: endpoints,
		Nodes:     nodes,
	}
}
