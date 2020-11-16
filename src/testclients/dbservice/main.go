package main

import (
	"context"
	"fmt"

	example "github.com/kennyzhu/go-os/src/dbservice/proto/example"
	"github.com/micro/go-micro"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := example.NewPreferencesService("go.micro.srv.dbservice", service.Client())

	// Make request
	rsp, err := cl.GetPreference(context.Background(), &example.PreferenceRequest{
		User: 1,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.ResultCode)
	fmt.Println(rsp.Prefer)

	rsp2, err := cl.GetPreferencesList(context.Background(), &example.PreferencesListRequest{
		Index: 1,
		Limit: 2,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp2.ResultCode)
	fmt.Println( len(rsp2.Prefers) )
	fmt.Println (rsp2.Prefers)
}
