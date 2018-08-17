/*
@Time : 2018/8/16 9:55
@Author : kenny zhu
@File : tables.go
@Software: GoLand
@Others: Sharding tables' operations
*/
package models

import (
	"github.com/go-xorm/xorm"
	"github.com/kennyzhu/go-os/dbservice/conf"
	"strconv"
)

//TableHash name
func tableNameHash(tableName string, suffixKey int64) string {
	if conf.OrmConf.TableHashValue == 0 {
		return tableName
	}

	return string(tableName + "_" + strconv.FormatInt(suffixKey%int64(conf.OrmConf.TableHashValue), 10))
}

func ShardGetTable(tableNameOrBean interface{}, suffixKey int64) *xorm.Session {
	tbName := orm.TableName(tableNameOrBean)
	return orm.Table(tableNameHash(tbName, suffixKey))
}

//engine.Table("user").Select("user.*, detail.*").
//    Join("INNER", "detail", "detail.user_id = user.id").
//    Where("user.name = ?", name).Limit(10, 0).
//    Find(&users)
//Get object by primary key id
func ShardGetById(id int64, obj interface{}) error {
	//// 复合主键的获取方法..
	//// has, errr := engine.Id(xorm.PK{1,2}).Get(user)..
	has, err := ShardGetTable(obj, id).Get(obj)
	if err != nil {
		return err
	}
	if !has {
		return ErrNotExist
	}
	return nil
}

//根据结构体中已有的非空数据来判断是否存在...
func ShardIsExist(id int64, obj interface{}) bool {
	has, _ := ShardGetTable(obj, id).Get(obj)
	return has
}

//Shard find limit objects.
//ShardGetTable support array.
func ShardFind(id int64, limit, start int, objs interface{}) error {
	return ShardGetTable(objs, id).Limit(limit, start).Find(objs)
}
