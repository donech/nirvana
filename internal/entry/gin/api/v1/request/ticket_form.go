package request

type TicketForm struct {
	Number string `json:"number" form:"number" binding:"required"`
	Period string `json:"period" form:"period" binding:"required"`
	Type   string `json:"type" form:"type" binding:"required"`
}
