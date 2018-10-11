package main

import (
	"log"
	"time"

	"github.com/kennyzhu/go-os/dbservice/tools/example"
	"github.com/micro/go-micro"

)


func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	example.RegisterExampleServiceHandler(service.Server(), example.NewSay())

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
