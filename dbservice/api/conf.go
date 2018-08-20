/*
@Time : 2018/8/20 16:08 
@Author : kenny zhu
@File : conf.go
@Software: GoLand
@Others:
*/
package main

import (
	"github.com/micro/go-config"
	"fmt"
	"github.com/kennyzhu/go-os/log"
	"github.com/micro/go-config/source/file"
)

var AppConf struct{
	//logger..
	LogLevel int32 `json:"LogLevel"`
	LogPath  string    `json:"LogPath"`

	//srv set..
	ApiName string     `json:"ApiName"`
	SrvName string     `json:"SrvName"`
	VersionTag int     `json:"Version"`

	//Todo:
}

//self init...
func init() {
	//read config
	if err := config.Load(file.NewSource(
		file.WithPath("./conf/api.json"),
	)); err != nil {
		fmt.Println(err)
		return
	}

	if err := config.Get("api").Scan(&AppConf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(AppConf.ApiName)
	if AppConf.LogPath == "" {
		AppConf.LogPath = "/dev/stdout"
	}
	//init logger..
	log.InitLogger(
		log.WithLevel( log.Level(AppConf.LogLevel) ),
		log.WithFields(log.Fields{
			"logger": "api",
		}),
		log.WithOutput(
			log.NewOutput(log.OutputName(AppConf.LogPath)),
		),
	)

	log.Infof("logger init, path:%v", AppConf.LogPath)
}
