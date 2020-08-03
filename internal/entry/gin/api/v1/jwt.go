package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/donech/tool/entry/gin/middleware"
	"github.com/donech/tool/xjwt"
	"github.com/gin-gonic/gin"
)

type JwtController struct {
	jwtMiddleware *middleware.JWTMiddleware
}

func (j JwtController) Login() xjwt.LoginFunc {
	return func(form xjwt.LoginInForm) (jwt.MapClaims, error) {
		return jwt.MapClaims{"username": "solar", "password": "123456"}, nil
	}
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
