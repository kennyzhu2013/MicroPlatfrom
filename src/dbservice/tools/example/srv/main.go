package main

import (
	example2 "github.com/kennyzhu/go-os/src/dbservice/tools/example"
	"log"
	"time"

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
	example2.RegisterExampleServiceHandler(service.Server(), example2.NewExampleServiceServer())

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
