package api

import (
	"github.com/donech/nirvana/internal/code"
	"github.com/donech/nirvana/internal/common"
	"github.com/donech/nirvana/internal/domain/lottery/entity"
	"github.com/donech/nirvana/internal/domain/lottery/service"
	"github.com/donech/tool/entry/gin/middleware"
	"github.com/donech/tool/xlog"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func NewLotteryController(lotteryService *service.LotteryService, jwtMiddleware *middleware.JWTMiddleware) *LotteryController {
	return &LotteryController{lotteryService: lotteryService, jwtMiddleware: jwtMiddleware}
}

type LotteryController struct {
	lotteryService *service.LotteryService
	jwtMiddleware  *middleware.JWTMiddleware
}

func (c LotteryController) GetTicket(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt64()
	userId := common.GetUserID(ctx.Request.Context())
	xlog.S(ctx.Request.Context()).Info("userID is: ", userId)
	ticket, err := c.lotteryService.TicketByID(ctx.Request.Context(), id)
	if err != nil {
		xlog.S(ctx.Request.Context()).Info("get ticket failed :", err.Error())
	}
	ResponseJSON(ctx, code.Success, "success", ticket)
}

func (c LotteryController) CheckTicket(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt64()
	ticket, err := c.lotteryService.TicketCheck(ctx.Request.Context(), id)
	if err != nil {
		xlog.S(ctx.Request.Context()).Info("check ticket failed :", err.Error())
	}
	ResponseJSON(ctx, code.Success, "success", ticket)
}

type TicketForm struct {
	Number string `json:"number" form:"number" binding:"required"`
	Period string `json:"period" form:"period" binding:"required"`
	Type   string `json:"type" form:"type" binding:"required"`
}

func (c LotteryController) CreateTicket(ctx *gin.Context) {
	var ticketForm TicketForm
	err := ctx.ShouldBind(&ticketForm)
	if err != nil {
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	ticket := entity.LotteryTicket{
		UserID:     0,
		Number:     ticketForm.Number,
		TicketType: ticketForm.Type,
		Period:     ticketForm.Period,
	}
	err = c.lotteryService.CreateTicket(ctx.Request.Context(), &ticket)
	if err != nil {
		xlog.S(ctx.Request.Context()).Info("create ticket failed :", err.Error())
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	ResponseJSON(ctx, code.Success, "success", ticket)
}

type GetRecordForm struct {
	Period string `json:"period" form:"period" binding:"required"`
	Type   string `json:"type" form:"type" binding:"required,oneof=TwoToneSphere SuperLotto"`
}

func (c LotteryController) GetRecord(ctx *gin.Context) {
	var form GetRecordForm
	err := ctx.ShouldBindUri(&form)
	if err != nil {
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	record, err := c.lotteryService.RecordByPeriodAndType(ctx.Request.Context(), form.Period, form.Type)
	if err != nil {
		xlog.S(ctx.Request.Context()).Info("get record failed :", err.Error())
	}
	ResponseJSON(ctx, code.Success, "success", record)
}

func (c LotteryController) CreateRecord(ctx *gin.Context) {
	var form GetRecordForm
	err := ctx.ShouldBind(&form)
	if err != nil {
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	record, err := c.lotteryService.GenerateRecordByPeriodAndType(ctx.Request.Context(), form.Period, form.Type)
	if err != nil {
		xlog.S(ctx.Request.Context()).Info("get record failed :", err.Error())
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	ResponseJSON(ctx, code.Success, "success", record)
}

// @create lottery tables
// @Summary init  lottery table
// @Produce  json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /tool/migration/ticket [get]
func (c LotteryController) Migration(ctx *gin.Context) {
	c.lotteryService.Migration()
	ResponseJSON(ctx, code.Success, "", gin.H{"status": "success"})
}

func (c LotteryController) RegisterRoute(engine *gin.RouterGroup) {
	r := engine.Group("/v1/lottery").Use(c.jwtMiddleware.MiddleWareImpl())
	r.GET("/ticket/:id", c.GetTicket)
	r.POST("/ticket", c.CreateTicket)
	r.GET("/check/:id", c.CheckTicket)
	r.GET("/record/:type/:period", c.GetRecord)
	r.POST("/record", c.CreateRecord)
	t := engine.Group("/tool")
	t.GET("/migration/ticket", c.Migration)
}
