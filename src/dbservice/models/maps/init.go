/*
@Time : 2018/9/29 11:15 
@Author : kenny zhu
@File : init.go
@Software: GoLand
@Others:
*/


package maps

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var (
	// error code define here...
	ErrNotFound = errors.New("not found record")

	// singleton instance..
	orm *xorm.Engine
)

func Init(db *xorm.Engine) {
	orm = db
}

func SyncTables()  error {
	return orm.Sync2( new(Article), new(Category), )
}