package main

import (
	"log"
	"github.com/micro/go-web"

	. "github.com/kennyzhu/go-os/dbservice/api/gin/modules"
	"github.com/kennyzhu/go-os/dbservice/api/conf"
)

func main() {
	conf.Init( "./conf/ginapi.json" )

	// Create service
	// must micro api --handler=http
	service := web.NewService(
		web.Name(conf.AppConf.ApiName), // " go.micro.web.dbservice"
	)
	service.Init()

	// Register modules and app.Run...
	// All path processed by modules..
	service.Handle("/", Modules.Router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
