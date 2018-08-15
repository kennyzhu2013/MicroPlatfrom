package models

import (
	"os"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"

	//self library..
	"github.com/kennyzhu/go-os/dbservice/conf"
)

var (
	ErrNotExist = errors.New("not exist")

	//singleton instance..
    orm *xorm.Engine
)

//need test promode ..
func Init(isProMode bool) {
	var err error
	orm, err = xorm.NewEngine(conf.OrmConf.DriverName, conf.OrmConf.DataSource)
	if err != nil {
		panic(err)
	}

	orm.SetMaxIdleConns(conf.OrmConf.MaxIdle)
	orm.SetMaxOpenConns(conf.OrmConf.MaxOpen)
	if conf.OrmConf.DebugLog {
		orm.Logger().SetLevel(core.LOG_DEBUG)
	} else {
		orm.Logger().SetLevel(core.LOG_INFO)
	}


	if !isProMode {
		orm.ShowSQL(true)

		//simple log..
		sqlWriter,_ := os.Create(conf.LogOutPutPath + "sql.log")
		logger := xorm.NewSimpleLogger(sqlWriter)
		logger.ShowSQL(true)
		orm.SetLogger(logger)
	}

	//struct sync...
	//if not define then create...
	/*err = orm.Sync2(new(Setting), new(Category), new(Post), new(Image),
		new(User), new(FavoritePost), new(Follow), new(Topic), new(FollowTopic),
		new(Page), new(Notification), new(Comment), new(Bulletin))
	if err != nil {
		panic(err)
	}*/
	err = orm.Sync2(new(Preferences))
	if err != nil {
		panic(err)
	}

	//orm set...
	//social.SetORM(orm)
}
