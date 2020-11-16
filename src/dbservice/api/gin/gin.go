package main

import (
	conf2 "github.com/kennyzhu/go-os/src/dbservice/api/conf"
	"github.com/kennyzhu/go-os/src/dbservice/api/gin/modules"
	"github.com/micro/go-web"
	"log"
)

func main() {
	conf2.Init( "./conf/ginapi.json" )

	// Create service
	// must micro api --handler=http
	service := web.NewService(
		web.Name(conf2.AppConf.ApiName), // " go.micro.web.dbservice"
	)
	service.Init()

	// Register modules and app.Run...
	// All path processed by modules..
	service.Handle("/", modules.Modules.Router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
