package entity

import "github.com/donech/tool/xdb"

type User struct {
	xdb.Entity
	Name  string
	Phone string
}

func (e User) TableName() string {
	return "user"
}
