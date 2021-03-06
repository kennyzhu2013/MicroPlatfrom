package models

import (
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"

	// self library..
	. "github.com/kennyzhu/go-os/src/dbservicevice/conf"
	"github.com/kennyzhu/go-os/src/dbservicevice/models/maps"
)

var (
	ErrNotExist = errors.New("not exist")

	// singleton instance..
    orm *xorm.Engine
)

// tables synced..
func SyncAllTables() error {
	// struct sync...
	// if not define then create...
	/*err = orm.Sync2(new(Setting), new(Category), new(Post), new(Image),
		new(User), new(FavoritePost), new(Follow), new(Topic), new(FollowTopic),
		new(Page), new(Notification), new(Comment), new(Bulletin))
	if err != nil {
		panic(err)
	}*/
	err := maps.SyncTables()
	if err != nil {
		return err
	}
	return orm.Sync2(new(Preferences), )
}

// need test pro-mode ..
func Init(isProMode bool, logPath string) {
	var err error
	orm, err = xorm.NewEngine(OrmConf.DriverName, OrmConf.DataSource)
	if err != nil {
		panic(err)
	}

	err = orm.Ping()
	if err != nil {
		panic(err)
	}

	orm.SetMaxIdleConns(OrmConf.MaxIdle)
	orm.SetMaxOpenConns(OrmConf.MaxOpen)
	if OrmConf.DebugLog {
		orm.Logger().SetLevel(core.LOG_DEBUG)
	} else {
		orm.Logger().SetLevel(core.LOG_INFO)
	}

	// for debug mode
	if !isProMode {
		orm.ShowSQL(true)

		// simple log..
		if logPath == "" {
			logPath = "./sql.log"
		}
		sqlWriter,_ := os.Create(logPath)
		logger := xorm.NewSimpleLogger(sqlWriter)
		logger.ShowSQL(true)
		orm.SetLogger(logger)
	}

	// sql cache for orm query, results count.
	if OrmConf.IsCached {
		ormCache := xorm.NewLRUCacher2(xorm.NewMemoryStore(), OrmConf.CacheTime, OrmConf.CacheCount)
		orm.SetDefaultCacher(ormCache)
	}

	// sql map for orm.
	err = orm.RegisterSqlMap( xorm.Xml(SqlMap.XmlLocation, SqlMap.XmlSuffix) )
	if err != nil {
		panic(err)
	}
	if true == SqlMap.OpenWatcher {
		err = orm.StartFSWatcher()
		if err != nil {
			panic(err)
		}
	}

	// init models and sync tables...
	maps.Init( orm )
	err = SyncAllTables()
	if err != nil {
		panic(err)
	}

	fmt.Println("Models Init successed...")
	// orm set...
	// social.SetORM(orm)
}
