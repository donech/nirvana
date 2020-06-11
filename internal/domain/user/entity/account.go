package entity

import "github.com/donech/tool/xdb"

type Account struct {
	xdb.Entity
	Account  string `json:"account"`
	Password string `json:"password"`
	Token    string `json:"token"`
	xdb.CUDTime
}

func (e Account) TableName() string {
	return "account"
}
