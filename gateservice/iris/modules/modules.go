/*
@Time : 2018/8/21 11:05 
@Author : kenny zhu
@File : modules.go
@Software: GoLand
@Others:
*/
package modules

import (
	"github.com/kataras/iris"

	"github.com/kennyzhu/go-os/log"
)

var Modules struct{
	App *iris.Application

}

// self init...
func init() {
	Modules.App = iris.Default()
}

func NoModules(ctx iris.Context) {
	log.Info("Received NoModules API request")

	ctx.JSON(iris.Map{
		"message": "No module defined!",
	})
}

//anything to do..
//run here not go routine...