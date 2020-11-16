/*
@Time : 2018/8/21 11:28 
@Author : kenny zhu
@File : json.go
@Software: GoLand
@Others:
*/
package conf

import (
	"crypto/tls"
	"fmt"
	log2 "github.com/kennyzhu/go-os/src/log"
	"github.com/micro/go-config"
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
	log2.InitLogger(
		log2.WithLevel( log2.Level(AppConf.LogLevel) ),
		log2.WithFields(log2.Fields{
			"logger": "api",
		}),
		log2.WithOutput(
			log2.NewOutput(log2.OutputName(AppConf.LogPath)),
		),
	)

	log2.Infof("logger init, path:%v", AppConf.LogPath)
}

