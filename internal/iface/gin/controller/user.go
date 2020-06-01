package controller

import (
	"strconv"

	"github.com/prometheus/common/log"

	"github.com/donech/nirvana/internal/domain/user/service"
	"github.com/gin-gonic/gin"
)

func NewUserController(userSimpleService *service.SimpleService) *UserController {
	return &UserController{UserSimpleService: userSimpleService}
}

type UserController struct {
	UserSimpleService *service.SimpleService
}

func (c UserController) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			log.Fatalln("id must convert failed", err.Error())
		}
		user := c.UserSimpleService.ItemByID(id)
		ctx.JSON(200, gin.H{"user": user})
	}
}

func (c UserController) RegisterRoute(engine *gin.Engine) {
	user := engine.Group("/user")
	user.GET("/:id", c.GetUser())
}
