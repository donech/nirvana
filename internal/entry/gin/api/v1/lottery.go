package v1

import (
	"github.com/donech/core/xlog"
	"github.com/donech/nirvana/internal/code"
	"github.com/donech/nirvana/internal/domain/lottery/entity"
	"github.com/donech/nirvana/internal/domain/lottery/service"
	"github.com/donech/nirvana/internal/entry/gin/api/v1/request"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func NewLotteryController(lotteryService *service.LotteryService) *LotteryController {
	return &LotteryController{lotteryService: lotteryService}
}

type LotteryController struct {
	lotteryService *service.LotteryService
}

func (c LotteryController) GetTicket(ctx *gin.Context) {
	id := com.StrTo(ctx.Param("id")).MustInt64()
	ticket, err := c.lotteryService.TicketByID(ctx.Request.Context(), id)
	if err != nil {
		xlog.Ctx(ctx.Request.Context()).Info("get ticket failed :", err.Error())
	}
	ResponseJSON(ctx, code.Success, "success", ticket)
}

func (c LotteryController) CreateTicket(ctx *gin.Context) {
	var ticketForm request.TicketForm
	err := ctx.ShouldBind(&ticketForm)
	if err != nil {
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	err = c.lotteryService.CreateTicket(ctx.Request.Context(), entity.LotteryTicket{
		UserID:     0,
		Number:     ticketForm.Number,
		TicketType: ticketForm.Type,
		Period:     ticketForm.Period,
	})
	if err != nil {
		xlog.Ctx(ctx.Request.Context()).Info("create ticket failed :", err.Error())
		ResponseJSON(ctx, code.Error, err.Error(), nil)
		return
	}
	ResponseJSON(ctx, code.Success, "success", nil)
}

// @创建用户表
// @Summary init user table
// @Produce  json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /tool/migration/user [get]
func (c LotteryController) Migration(ctx *gin.Context) {
	c.lotteryService.Migration()
	ResponseJSON(ctx, code.Success, "", gin.H{"status": "success"})
}

func (c LotteryController) RegisterRoute(engine *gin.RouterGroup) {
	r := engine.Group("/v1/ticket")
	r.GET("/:id", c.GetTicket)
	r.POST("/", c.CreateTicket)
	t := engine.Group("/tool")
	t.GET("/migration/ticket", c.Migration)
}
