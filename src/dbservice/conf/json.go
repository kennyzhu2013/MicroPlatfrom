package conf

import (
	//"encoding/json"
	//"io/ioutil"
	"fmt"
	log2 "github.com/kennyzhu/go-os/src/log"

	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
	"time"
)

// struct define...
// var cc = make(map[string]string) //cache map
// OrmConf config as variable..
var OrmConf struct{
	DriverName string `json:"DriverName"`
	DataSource string `json:"DataSource"`
	MaxIdle    int    `json:"MaxIdle"`
	MaxOpen    int    `json:"MaxOpen"`
	DebugLog   bool   `json:"DebugLog"`
	IsProMode  bool   `json:"IsProMode"`

	// cache
	IsCached  			bool   `json:"UseCache"`
	CacheTime time.Duration `json:"CacheDuration"`
	CacheCount 			int `json:"CacheRecordCount"`

	// if TableHashValue == 0, no table sharding..
	TableHashValue		int `json:"TableHashValue"`

	//Todo:
	DatabaseHashValue 	int `json:"DbHashValue"`
}

var SqlMap struct{
	OpenWatcher bool `json:"FSWatcher"`
	XmlLocation string `json:"MapLocation"`
	XmlSuffix string `json:"MapSuffix"`

	// ?...stpl
}

var AppConf struct{
	//logger..
	LogLevel int32 `json:"LogLevel"`
	LogPath  string    `json:"LogPath"`

	// srv set..
	SrvName string     `json:"SrvName"`
	VersionTag int     `json:"Version"`

	// Todo:
	SetId             	int `json:"GlobalSetId"`
}

// self init...
func init() {
	// read config
	if err := config.Load(file.NewSource(
		file.WithPath("./conf/dbserver.json"),
	)); err != nil {
		fmt.Println(err)
		return
	}

	// read OrmConf...
	if err := config.Get("database").Scan(&OrmConf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(OrmConf.DriverName, OrmConf.DataSource)
	if OrmConf.MaxIdle == 0 { OrmConf.MaxIdle = OrmMaxIdleConn
	}
	if OrmConf.MaxOpen == 0 { OrmConf.MaxOpen = OrmMaxOpen
	}

	if OrmConf.IsCached == true {
		if OrmConf.CacheTime == 0 { OrmConf.CacheTime = OrmCacheTime
		}
		if OrmConf.CacheCount == 0 { OrmConf.CacheCount = OrmCacheCount
		}
	}

	// load all sql map
	if err := config.Get("sqlmap").Scan(&SqlMap); err != nil {
		fmt.Println(err)
		return
	}
	if SqlMap.XmlSuffix == "" {
		SqlMap.XmlSuffix = ".xml"
	}

	if err := config.Get("app").Scan(&AppConf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(AppConf.SrvName)
	if AppConf.LogPath == "" {
		AppConf.LogPath = LogOutPutPath
	}
	// init logger..
	log2.InitLogger(
		log2.WithLevel( log2.Level(AppConf.LogLevel) ),
		log2.WithFields(log2.Fields{
			"logger": LoggerFlag,
		}),
		log2.WithOutput(
			log2.NewOutput(log2.OutputName(AppConf.LogPath)),
		),
	)
	log2.Infof("logger init, path:%v", AppConf.LogPath)
}
