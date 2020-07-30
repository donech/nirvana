package v1

import (
	_ "github.com/donech/nirvana/internal/entry/gin/docs"
	"github.com/donech/tool/xlog/ginzap"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewDefaultController() *DefaultController {
	return &DefaultController{}
}

type DefaultController struct {
}

func (c DefaultController) RegisterRoute(root *gin.RouterGroup) {
	root.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "health", "connection_num": ginzap.GetConnectionNum(),
		})
	})
	root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
