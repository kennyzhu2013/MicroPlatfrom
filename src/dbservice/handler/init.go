/*
@Time : 2018/8/16 15:26 
@Author : kenny zhu
@File : init.go
@Software: GoLand
@Others:
*/
package handler

import (
	"github.com/kennyzhu/go-os/src/dbservice/proto/example"
	"github.com/micro/go-micro/server"
)

//all handlers init here...
func Init(s server.Server) {
	//register handler here, each request run in goroutine...
	//example.RegisterPreferencesHandler(service.Server(), new(handler.Example))
	go_micro_srv_dbservice.RegisterPreferencesHandler(s, new(Example))

	//To register others...
}