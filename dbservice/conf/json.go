package conf

import (
	//"encoding/json"
	//"io/ioutil"
	"fmt"

	"github.com/kennyzhu/go-os/log"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	//"github.com/kennyzhu/go-os/config/source/file"
)

//struct define...
//var cc = make(map[string]string) //cache map
//OrmConf config as variable..
var OrmConf struct{
	DriverName string `json:"DriverName"`
	DataSource string `json:"DataSource"`
	MaxIdle    int    `json:"MaxIdle"`
	MaxOpen    int    `json:"MaxOpen"`
	DebugLog   bool   `json:"DebugLog"`
	IsProMode  bool   `json:"IsProMode"`
}

var AppConf struct{
	//logger..
	LogLevel log.Level `json:"LogLevel"`
	LogPath  string    `json:"LogPath"`

	//srv set..
	SrvName string     `json:"SrvName"`
	VersionTag int     `json:"Version"`
}

//self init...
func init() {
	//read config
	if err := config.Load(file.NewSource(
		file.WithPath("./bin/dbserver.json"),
	)); err != nil {
		fmt.Println(err)
		return
	}

	//read OrmConf...
	if err := config.Get("database").Scan(&OrmConf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(OrmConf.DriverName, OrmConf.DataSource)
	if OrmConf.MaxIdle == 0 {
		OrmConf.MaxIdle = OrmMaxIdleConn
	}
	if OrmConf.MaxOpen == 0 {
		OrmConf.MaxOpen = OrmMaxOpen
	}

	if err := config.Get("app").Scan(&AppConf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(AppConf.SrvName)
	if AppConf.LogPath == "" {
		AppConf.LogPath = LogOutPutPath
	}
	//init logger..
	logger := log.NewLog(
		log.WithLevel(AppConf.LogLevel),
		log.WithFields(log.Fields{
			"logger": LoggerFlag,
		}),
		log.WithOutput(
			log.NewOutput(log.OutputName(AppConf.LogPath)),
		),
	)
	logger.Debug("logger init, path:%v", AppConf.LogPath)
}
