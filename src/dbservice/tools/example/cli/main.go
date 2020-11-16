/*
@Time : 2018/10/10 16:41 
@Author : kenny zhu
@File : main.go
@Software: GoLand
@Others:
*/
package main

import (
	"context"
	"fmt"
	example2 "github.com/kennyzhu/go-os/src/dbservice/tools/example"
	"github.com/micro/go-micro"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := example2.NewExampleService("go.micro.srv.greeter", service.Client())

	// Make request
	rsp, err := cl.GetPhone(context.Background(), &example2.GetPhoneRequest{
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Output is:")
	fmt.Println(rsp.Name)
	fmt.Println(rsp.Price)
	fmt.Println(rsp.CategoryID)
}