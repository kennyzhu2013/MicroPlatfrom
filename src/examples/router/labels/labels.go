package main

import (
	"fmt"

	hello "github.com/micro/examples/greeter/srv/proto/hello"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-os/router"

	"golang.org/x/net/context"
)

func main() {
	cmd.Init()

	r := router.NewRouter()
	c := client.NewClient(client.Selector(r))
	c = router.NewLabelWrapper(c)

	// Create new request to service go.micro.srv.greeter, method Say.Hello
	req := c.NewRequest("go.micro.srv.greeter", "Say.Hello", &hello.Request{
		Name: "John",
	})

	rsp := &hello.Response{}

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		router.LabelPrefix + "Greeter": "one",
	})

	// Call service
	if err := c.Call(ctx, req, rsp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
