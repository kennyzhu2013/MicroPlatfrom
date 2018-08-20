/*
@Time : 2018/8/15 19:05 
@Author : kenny zhu
@File : base.go
@Software: GoLand
@Others: model base operation
@refer：github.com/go-tango/wego/blob/master/models/base.go.
*/
package models

import (
	"github.com/go-xorm/xorm"
	"bytes"
)

func ORM() *xorm.Engine {
	return orm
}

type ModelFinder interface {
	Object() interface{}
}

//Get object by primary key id
func GetById(id int64, obj interface{}) error {
	//// 复合主键的获取方法..
	//// has, errr := engine.Id(xorm.PK{1,2}).Get(user)..
	//fmt.Printf("id:%d", id)
	//fmt.Printf("obj type:%v", reflect.TypeOf(obj))
	has, err := orm.Id(id).Get(obj)
	if err != nil {
		return err
	}
	if !has {
		return ErrNotExist
	}
	return nil
}

//根据结构体中已有的非空数据来判断是否存在...
func IsExist(obj interface{}) bool {
	has, _ := orm.Get(obj)
	return has
}

//find limit objects.
//objs为slice的指针或Map指针，即为查询后返回的结果...
func Find(limit, start int, objs interface{}) error {
	return orm.Limit(limit, start).Find(objs)
}

//add other ops for all models..
// Create tables for QReader, this will drop them first, you may lost data if the tables are already exists.
func CreateTables(createTablesSql string) error {
	sql := "begin;" + createTablesSql + "commit;"
	_, err := orm.Import(bytes.NewReader([]byte(sql)))
	return err
}

// Create indexes.
func CreateIndexes(createIndexesSql string) error {
	sql := "begin;" + createIndexesSql + "commit;"
	_, err := orm.Import(bytes.NewReader([]byte(sql)))
	return err
}


