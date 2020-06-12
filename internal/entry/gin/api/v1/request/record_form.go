package request

type GetRecordForm struct {
	Period string `json:"period" form:"period" binding:"required"`
	Type   string `json:"type" form:"type" binding:"required"`
}
