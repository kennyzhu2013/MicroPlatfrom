/*
@Time : 2018/8/21 11:05 
@Author : kenny zhu
@File : modules.go
@Software: GoLand
@Others:
*/
package modules

import (

	"github.com/kennyzhu/go-os/log"
	"github.com/gin-gonic/gin"
)

var Modules struct{
}

func NoModules(ctx *gin.Context) {
	log.Info("Received NoModules API request")

	ctx.JSON(200, map[string]string{
		"message": "No module defined!",
	})
}

//anything to do..
//run here not go routine...