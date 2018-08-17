package models

import (
	"os"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"

	//self library..
	. "github.com/kennyzhu/go-os/dbservice/conf"
)

var (
	ErrNotExist = errors.New("not exist")

	//singleton instance..
    orm *xorm.Engine
)

func SyncAllTables() error {
	//struct sync...
	//if not define then create...
	/*err = orm.Sync2(new(Setting), new(Category), new(Post), new(Image),
		new(User), new(FavoritePost), new(Follow), new(Topic), new(FollowTopic),
		new(Page), new(Notification), new(Comment), new(Bulletin))
	if err != nil {
		panic(err)
	}*/
	return orm.Sync2(new(Preferences), )
}

//need test promode ..
func Init(isProMode bool) {
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


	if !isProMode {
		orm.ShowSQL(true)

		//simple log..
		sqlWriter,_ := os.Create("./log/sql.log")
		logger := xorm.NewSimpleLogger(sqlWriter)
		logger.ShowSQL(true)
		orm.SetLogger(logger)
	}

	if OrmConf.IsCached {
		ormCache := xorm.NewLRUCacher2(xorm.NewMemoryStore(), OrmConf.CacheTime, OrmConf.CacheCount)
		orm.SetDefaultCacher(ormCache)
	}

	err = SyncAllTables()
	if err != nil {
		panic(err)
	}

	fmt.Println("Models Init successed...")
	//orm set...
	//social.SetORM(orm)
}
