package api

import (
	"github.com/donech/tool/entry/gin/middleware"
	"github.com/gin-gonic/gin"
)

type JwtController struct {
	jwtMiddleware *middleware.JWTMiddleware
}

func (j JwtController) TestJwt(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "success"})
}

func (j JwtController) RegisterRoute(engine *gin.RouterGroup) {
	engine.GET("token", j.jwtMiddleware.GenerateTokenHandler())
	auth := engine.Group("auth")
	auth.Use(j.jwtMiddleware.MiddleWareImpl())
	auth.GET("test", j.TestJwt)
}

func NewJwtController(jwt *middleware.JWTMiddleware) *JwtController {
	return &JwtController{jwtMiddleware: jwt}
}
