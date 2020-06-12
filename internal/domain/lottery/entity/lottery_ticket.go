package entity

import "github.com/donech/tool/xdb"

type LotteryTicket struct {
	xdb.Entity
	UserID     int64  `json:"user_id"`
	Number     string `json:"number"`
	TicketType string `json:"ticket_type"`
	Period     string `json:"period"`
	Level      int    `json:"level"`
	Price      int    `json:"price"`
	xdb.CUDTime
}

func (t LotteryTicket) TableName() string {
	return "lottery_ticket"
}
