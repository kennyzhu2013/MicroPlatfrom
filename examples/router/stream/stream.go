package main

import (
	"fmt"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"

	proto "github.com/micro/router-srv/proto/router"

	"golang.org/x/net/context"
)

var (
	service = "go.micro.srv.router"
)

func main() {
	cmd.Init()

	r := proto.RouterServiceClient(service, client.DefaultClient)

	stream, err := r.SelectStream(context.TODO(), &proto.SelectRequest{Service: service})
	if err != nil {
		fmt.Println("error streaming", err)
		return
	}

	for i := 0; i <= 3; {
		fmt.Println("waiting on stream")
		rsp, err := stream.Recv()
		if err != nil {
			fmt.Println("error receiving", err)
			return
		}
		fmt.Println("got stream response, expires", rsp.Expires)
		for _, s := range rsp.Services {
			fmt.Printf("received %s %s %+v\n", s.Name, s.Version, s.Nodes)
		}
	}
}
