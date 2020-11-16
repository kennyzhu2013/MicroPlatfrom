/*
@Time : 2018/8/21 11:04 
@Author : kenny zhu
@File : init.go
@Software: GoLand
@Others:
*/
package modules

import (
	conf2 "github.com/kennyzhu/go-os/src/dbservice/api/conf"
	"github.com/kennyzhu/go-os/src/dbservice/proto/example"
	"github.com/micro/go-micro/client"
)

// All handlers init here...
func Init() {
	// default ..
	Modules.App.Post("/", NoModules)
	Modules.App.Get("/", NoModules)

	// Examples Begin
	// Base module router for rest full, Preferences is name of table or tables while Module equals database.
	e := new(examples)
	// Init Preferences Server Client
	e.cl = go_micro_srv_dbservice.NewPreferencesService(conf2.AppConf.SrvName, client.DefaultClient)
	Modules.App.Get("/Preferences/{action:path}", e.Preferences)
	// Examples End

	// register other handlers here, each request run in goroutine.

	// To register others...
}
