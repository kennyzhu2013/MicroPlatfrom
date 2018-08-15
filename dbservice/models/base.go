/*
@Time : 2018/8/15 19:05 
@Author : kenny zhu
@File : base.go
@Software: GoLand
@Others: model base operation
*/
package models

import (
	"github.com/go-xorm/xorm"
)

func ORM() *xorm.Engine {
	return orm
}

//Get object by primary key id
func GetById(id int64, obj interface{}) error {
	has, err := orm.Id(id).Get(obj)
	if err != nil {
		return err
	}
	if !has {
		return ErrNotExist
	}
	return nil
}

