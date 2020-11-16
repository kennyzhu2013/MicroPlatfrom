/*
@Time : 2018/9/26 16:21 
@Author : kenny zhu
@File : sample_impl.go
@Software: GoLand
@Others: 此文件需要用proteus工具生成对应的micro server文件...
*/
package maps

import (
	"time"
	"github.com/kennyzhu/vlcms/dao/models"
)

type JSONTime time.Time
type Article struct {
	Id             int       `xorm:"not null pk autoincr unique INTEGER"`
	Content        string    `xorm:"not null TEXT"`
	Title          string    `xorm:"not null VARCHAR(255)"`
	Categorysubid  int       `xorm:"not null INTEGER"`
	Remark         string    `xorm:"not null VARCHAR(2555)"`
	Userid         int       `xorm:"not null INTEGER"`
	Viewcount      int       `xorm:"not null default 0 INTEGER"`
	Replycount     int       `xorm:"not null default 0 INTEGER"`
	Tags           string    `xorm:"not null VARCHAR(300)"`
	Createdatetime JSONTime  `xorm:"not null default 'now()' DATETIME"`
	Isdraft        int       `xorm:"SMALLINT"`
	Lastupdatetime time.Time `xorm:"not null default 'now()' DATETIME"`
}

type Category struct {
	Id             int       `xorm:"not null pk autoincr unique INTEGER"`
	Name           string    `xorm:"not null VARCHAR(200)"`
	Counts         int       `xorm:"not null default 0 INTEGER"`
	Orders         int       `xorm:"not null default 0 INTEGER"`
	Createtime     time.Time `xorm:"not null default 'now()' created DATETIME"`
	Pid            int       `xorm:"not null default 0 INTEGER"`
	Lastupdatetime time.Time `xorm:"not null default 'now()' updated  DATETIME"`
	Status         int       `xorm:"not null default 1 SMALLINT"`
}

// service..
// paramMap := map[string]interface{}{"1": 2, "2": 5}..
func SelectAllArticle(paramMap map[string]interface{}) ([]Article,error) {
	var article []Article
	session := models.ORM().NewSession()
	result := session.SqlMapClient("selectAllArticle", &paramMap).Search(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	return article, nil
}

func GetAllCategories()  ([]Category,error) {
	var categories []Category
	session := models.ORM().NewSession()
	result := session.SqlMapClient("category").Search(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}


