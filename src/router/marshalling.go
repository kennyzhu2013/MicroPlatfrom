package router

import (
	"github.com/micro/go-micro/registry"
	router "github.com/micro/go-os/router/proto"
)

func values(v []*registry.Value) []*router.Value {
	if len(v) == 0 {
		return []*router.Value{}
	}

	var vs []*router.Value
	for _, vi := range v {
		vs = append(vs, &router.Value{
			Name:   vi.Name,
			Type:   vi.Type,
			Values: values(vi.Values),
		})
	}
	return vs
}

func toValues(v []*router.Value) []*registry.Value {
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

func serviceToProto(s *registry.Service) *router.Service {
	if s == nil || len(s.Nodes) == 0 {
		return nil
	}
	return &router.Service{
		Name:     s.Name,
		Version:  s.Version,
		Metadata: s.Metadata,
		Nodes: []*router.Node{&router.Node{
			Id:       s.Nodes[0].Id,
			Address:  s.Nodes[0].Address,
			Port:     int64(s.Nodes[0].Port),
			Metadata: s.Nodes[0].Metadata,
		}},
	}
}

func protoToService(s *router.Service) *registry.Service {
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
