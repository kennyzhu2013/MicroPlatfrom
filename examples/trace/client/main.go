package main

import (
	"fmt"
	example "github.com/micro/examples/server/proto/example"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-os/trace"
	"golang.org/x/net/context"
	"time"
)

func call(i int) {
	// Create new request to service go.micro.srv.example, method Example.Call
	req := client.NewRequest("go.micro.srv.example", "Ping.Pong", &example.Ping{
		Stroke: 10,
	})

	rsp := &example.Pong{}

	// Call service
	if err := client.Call(context.TODO(), req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("Call:", i, "rsp:", rsp.Stroke)
}

func main() {
	cmd.Init()

	t := trace.NewTrace()
	defer t.Close()

	srv := &registry.Service{
		Name: "go.micro.client.example",
	}

	client.DefaultClient = client.NewClient(
		client.Wrap(
			trace.ClientWrapper(t, srv),
		),
	)

	fmt.Println("\n--- Traced Call example ---")
	i := 0
	for {
		call(i)
		i++
		<-time.After(time.Second * 5)
	}
}
