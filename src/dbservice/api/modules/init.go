/*
@Time : 2018/8/21 11:04 
@Author : kenny zhu
@File : init.go
@Software: GoLand
@Others:
*/
package modules

import (
	"fmt"
	conf2 "github.com/kennyzhu/go-os/src/dbservice/api/conf"
	"github.com/kennyzhu/go-os/src/dbservice/proto/example"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
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
			&examples{cl: go_micro_srv_dbservice.NewPreferencesService(conf2.AppConf.SrvName, c)},
		),
	)

	// register other handlers here, each request run in goroutine.

	// To register others...

	fmt.Println("Modules init finished.")
}
