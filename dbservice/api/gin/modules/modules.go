/*
@Time : 2018/8/21 11:05 
@Author : kenny zhu
@File : modules.go
@Software: GoLand
@Others:
*/
package modules

import (
	"github.com/gin-gonic/gin"

	"github.com/kennyzhu/go-os/log"
)

var Modules struct{
	Router *gin.Engine
}

// self init...
func init() {
	//gin.SetMode(gin.ReleaseMode)
	Modules.Router = gin.Default()
	Init()
}

func NoModules(ctx *gin.Context) {
	log.Info("Received NoModules API request")

	ctx.JSON(200, map[string]string{
		"message": "No module defined!",
	})
}

//anything to do..
//run here not go routine...