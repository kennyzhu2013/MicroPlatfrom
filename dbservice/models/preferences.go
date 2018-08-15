//create by xorm reverse mysql command..
package models

type Preferences struct {
	User  int    `xorm:"not null pk INT(10)"`
	Name  string `xorm:"not null pk VARCHAR(32)"`
	Value string `xorm:"not null VARCHAR(64)"`
}
