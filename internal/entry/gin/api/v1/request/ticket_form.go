package request

type TicketForm struct {
	Number string `json:"number" binding:"required"`
	Period string `json:"period" binding:"required"`
	Type   string `json:"type" binding:"required"`
}
