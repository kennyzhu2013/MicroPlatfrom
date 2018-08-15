package main

import (
	"fmt"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-os/discovery"
	"time"
)

func main() {
	cmd.Init()
	d := discovery.NewDiscovery(
		discovery.Interval(time.Second),
	)

	defer d.Close()

	service := &registry.Service{
		Name:    "go.micro.srv.foo",
		Version: "latest",
		Metadata: map[string]string{
			"foo": "bar",
		},
		Endpoints: []*registry.Endpoint{
			&registry.Endpoint{
				Name: "/index",
				Request: &registry.Value{
					Name: "request",
					Type: "Request",
				},
				Response: &registry.Value{
					Name: "response",
					Type: "Response",
				},
				Metadata: map[string]string{
					"index": "Handles index requests",
				},
			},
		},
		Nodes: []*registry.Node{
			&registry.Node{
				Id:      "foo-123",
				Address: "localhost",
				Port:    8080,
				Metadata: map[string]string{
					"bar": "baz",
				},
			},
		},
	}

	// register
	if err := d.Register(service); err != nil {
		fmt.Println(err)
	}
	// find self
	rsp, err := d.GetService("go.micro.srv.foo")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Got service %+v\n", rsp[0])
	}

	<-time.After(time.Second * 10)

	// deregister
	if err := d.Deregister(service); err != nil {
		fmt.Println(err)
	}

}
