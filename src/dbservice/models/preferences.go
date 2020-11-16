/*
@Time : 2018/8/15 19:05
@Author : kennyzhu
@File : base.go
@Software: GoLand
@Others: create by xorm reverse mysql command.
*/
package models

type Preferences struct {
	User  int    `xorm:"not null pk autoincr INT(10)"`
	Name  string `xorm:"not null VARCHAR(32)"`
	Value string `xorm:"not null VARCHAR(64)"`

}
