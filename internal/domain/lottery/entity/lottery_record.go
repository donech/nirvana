package entity

import "github.com/donech/tool/xdb"

type LotteryRecord struct {
	xdb.Entity
	Number string `json:"number"`
	Period string `json:"period"`
	Type   string `json:"type"`
	xdb.CUDTime
}

func (e LotteryRecord) TableName() string {
	return "lottery_record"
}
