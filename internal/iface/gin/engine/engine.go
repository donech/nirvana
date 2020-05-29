package engine

import (
	"github.com/gin-gonic/gin"
)

//Engine *gin.Engine
var Engine *gin.Engine

func init() {
	Engine = gin.Default()
	Engine.GET("/redness", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"ping": "pong",
		})
	})
}
