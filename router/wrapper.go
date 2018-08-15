package router

import (
	"strings"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"

	"golang.org/x/net/context"
)

type labelWrapper struct {
	client.Client
}

func (l *labelWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)

	labels := make(map[string]string)

	for k, v := range md {
		if strings.HasPrefix(k, LabelPrefix) {
			key := strings.TrimPrefix(k, LabelPrefix)
			key = strings.ToLower(key)
			labels[key] = strings.ToLower(v)
		}
	}

	filter := func(services []*registry.Service) []*registry.Service {
		var filtered []*registry.Service

	SERVICE:
		for _, service := range services {
			// check labels for the service
			for lk, lv := range labels {
				// if we match then append the service and return
				if v := service.Metadata[lk]; v == lv {
					filtered = append(filtered, service)
					continue SERVICE
				}
			}

			var nodes []*registry.Node

		NODE:
			for _, node := range service.Nodes {
				// check labels against the node
				for lk, lv := range labels {
					// if it matches then append node
					if v := node.Metadata[lk]; v == lv {
						nodes = append(nodes, node)
						continue NODE
					}
				}
			}

			if len(nodes) > 0 {
				service.Nodes = nodes
				filtered = append(filtered, service)
			}
		}

		return filtered
	}

	callOptions := append(opts, client.WithSelectOption(
		selector.WithFilter(filter),
	))

	return l.Client.Call(ctx, req, rsp, callOptions...)
}

func NewLabelWrapper(c client.Client) client.Client {
	return &labelWrapper{c}
}
