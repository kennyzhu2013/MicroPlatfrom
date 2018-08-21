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
)

// All handlers init here...
func Init() {
	// default ..
	Modules.App.Get("/", NoModules)

	// Examples Begin
	// Base module router for rest full, Preferences is name of table or tables while Module equals database.
	e := new( examples )
	// Init Preferences Server Client
	e.cl = example.NewPreferencesService(conf.AppConf.SrvName, client.DefaultClient)
	Modules.App.Get("/Preferences/{action:path}", e.Preferences)
	// Examples End

	// register other handlers here, each request run in goroutine.

	// To register others...
}
