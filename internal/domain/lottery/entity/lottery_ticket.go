package entity

import "github.com/donech/tool/xdb"

const TicketCheckedStatus = 1

type LotteryTicket struct {
	xdb.Entity
	UserID     int64  `json:"user_id"`
	Number     string `json:"number"`
	TicketType string `json:"ticket_type"`
	Period     string `json:"period"`
	Level      int    `json:"level"`
	Price      int    `json:"price"`
	Status     int    `json:"status"`
	xdb.CUDTime
}

func (t LotteryTicket) TableName() string {
	return "lottery_ticket"
}

func (t LotteryTicket) IsChecked() bool {
	if t.Status == TicketCheckedStatus {
		return true
	}
	return false
}
