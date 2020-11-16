/*
@Time : 2018/8/21 10:55 
@Author : kenny zhu
@File : iris.go
@Software: GoLand
@Others:
*/
package main

import (
	"fmt"
	"github.com/kataras/iris"
	conf2 "github.com/kennyzhu/go-os/src/gateservice/conf"
	"github.com/kennyzhu/go-os/src/gateservice/iris/modules"
	"github.com/micro/go-micro/registry"
	"github.com/pborman/uuid"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)


var (
	service = &registry.Service{
		Name: "go.micro.srv.iris",
		Nodes: []*registry.Node{
			{
				Id:      "go.micro.srv.iris-" + uuid.NewUUID().String(),
				Address: "localhost",
				Port:    8400,
			},
		},
	}
)

// This is example for api gates..
// To delete....
func main() {
	conf2.Init( "./conf/irisapi.json" )
	service.Name = conf2.AppConf.ApiName

	// Create service
	/*
	webservice := web.NewService(
		web.Name(conf.AppConf.ApiName),
	)
	webservice.Init()*/

	modules.Init()
	address := conf2.AppConf.IpAddress + ":" + strconv.Itoa( int(conf2.AppConf.Port) )
	go modules.Modules.App.Run( iris.Addr( address ) ) // eg:":8400"
	// Register modules and app.Run...
	// All path processed by modules..
	// service.Handle("/", Modules.App)
	registry.Register(service)

	// 通过registry可以获得服务器的ip和端口等信息...
	// find self
	rsp, err := registry.GetService(service.Name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Got service %+v\n", rsp[0])
		fmt.Printf("Nodes info %+v\n", rsp[0].Nodes[0])
	}

	notify := make(chan os.Signal, 1)
	signal.Notify(notify, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-notify

	registry.Deregister(service)
}
