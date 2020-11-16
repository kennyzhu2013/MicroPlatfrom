package main

import (
	"fmt"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-os/auth"
	"golang.org/x/net/context"
)

func main() {

	// create an auth client
	a := auth.NewAuth(
		auth.Id("asim"),
		auth.Secret("foobar"),
	)

	// retreive a token
	t, err := a.Token()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Retreived Token %+v\n", t)

	ctx := auth.ContextWithToken(context.TODO(), t)

	// introspect a token from context
	t, err = a.Introspect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Introspected Token %+v\n", t)

	req := client.NewRequest("go.micro.example", "Example.Method", "request")

	// check if context/request is authorized
	t, err = a.Authorized(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Authorized Token %+v\n", t)
}
