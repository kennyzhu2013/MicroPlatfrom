/*
@Time : 2018/8/21 11:04 
@Author : kenny zhu
@File : init.go
@Software: GoLand
@Others:
*/
package modules

import (
	"github.com/micro/go-micro/client"
	example "github.com/kennyzhu/go-os/dbservice/proto/example"
	"github.com/kennyzhu/go-os/dbservice/api/conf"

	"fmt"
)

// All handlers init here...
func Init() {
	// default ..
	// micro health go.micro.api.gin call this function..
	Modules.Router.GET("/", NoModules)
	//Modules.Router.GET("/dbservice", NoModules)

	// Examples Begin
	// Base module router for rest full, Preferences is name of table or tables while Module equals database.
	// Call url:curl "http://localhost:8004/Preferences/GetPreference?user=1"
	e := new( examples )
	// Init Preferences Server Client
	e.cl = example.NewPreferencesService(conf.AppConf.SrvName, client.DefaultClient)
	Modules.Router.GET("/Preferences/*action", e.Preferences)
	// Examples End

	// register other handlers here, each request run in goroutine.

	// To register others...

	fmt.Println("Modules init finished.")
}
