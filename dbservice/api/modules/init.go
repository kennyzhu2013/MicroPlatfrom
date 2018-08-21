/*
@Time : 2018/8/21 11:04 
@Author : kenny zhu
@File : init.go
@Software: GoLand
@Others:
*/
package modules

import (
	example "github.com/kennyzhu/go-os/dbservice/proto/example"
	"github.com/kennyzhu/go-os/dbservice/api/conf"
	"fmt"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/client"
)

// All handlers init here...
func Init ( s server.Server, c client.Client ) {
	// default ..
	// micro health go.micro.api.gin call this function..
	//Modules.Router.GET("/dbservice", NoModules)

	// Examples Begin
	// examples is your modules for service
	s.Handle(
		s.NewHandler(
			&examples{cl: example.NewPreferencesService(conf.AppConf.SrvName, c)},
		),
	)

	// register other handlers here, each request run in goroutine.

	// To register others...

	fmt.Println("Modules init finished.")
}
