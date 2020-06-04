package entity

import "github.com/donech/tool/xdb"

type User struct {
	xdb.Entity
	Name  string `json:"name"`
	Phone string `json:"phone"`
	xdb.CUDTime
}

func (e User) TableName() string {
	return "user"
}
