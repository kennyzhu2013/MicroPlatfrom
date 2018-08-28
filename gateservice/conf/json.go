/*
@Time : 2018/8/21 11:28 
@Author : kenny zhu
@File : json.go
@Software: GoLand
@Others:
*/
package conf

import (
	"github.com/micro/go-config"
	"fmt"
	"github.com/kennyzhu/go-os/log"
	"github.com/micro/go-config/source/file"
	"crypto/tls"
)

var AppConf struct{
	//logger..
	LogLevel int32 `json:"LogLevel"`
	LogPath  string    `json:"LogPath"`

	//srv set..
	ApiName string     `json:"ApiName"`
	SrvName string     `json:"SrvName"`
	VersionTag int     `json:"Version"`

	// IP : port
	IpAddress string    `json:"IpAddress"`
	Port 	  int32     `json:"Port"`

	//Todo:
}

// self init...configFile = "./conf/api.json"...
func Init(configFile string) {
	// read config
	if err := config.Load(file.NewSource(
		file.WithPath(configFile),
	)); err != nil {
		fmt.Println(err)
		return
	}

	// Todo: tls config set...
	TlsCfg = new(tls.Config)

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

