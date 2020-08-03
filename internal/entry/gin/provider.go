package gin

import (
	"log"

	"github.com/dgrijalva/jwt-go"

	"github.com/donech/tool/entry/gin/middleware"

	"github.com/donech/nirvana/internal/config"
	v1 "github.com/donech/nirvana/internal/entry/gin/api/v1"
	_ "github.com/donech/nirvana/internal/entry/gin/docs"
	"github.com/donech/tool/entry/gin"
	"github.com/donech/tool/xjwt"

	"github.com/google/wire"
	"go.uber.org/zap"
)

func NewRouter(u *v1.UserController, l *v1.LotteryController, d *v1.DefaultController, jwt *v1.JwtController) gin.Router {
	r := &gin.DefaultRouter{}
	r.RegisterController(u)
	r.RegisterController(l)
	r.RegisterController(d)
	r.RegisterController(jwt)
	return r
}

func NewEntry(conf *config.Config, router gin.Router, logger *zap.Logger) *gin.Entry {
	confG := &gin.Config{
		Mod:  conf.Application.Mod,
		Addr: conf.Gin.Addr,
	}
	return gin.NewEntry(confG, router, logger)
}

func NewJWTFactory(conf *config.Config, loginFunc xjwt.LoginFunc) xjwt.JWTFactory {
	f, err := xjwt.NewJWTFactory(conf.JWT, xjwt.WithLoginFunc(loginFunc))
	if err != nil {
		log.Fatal(err.Error())
	}
	return f
}

func NewLoginFunc() xjwt.LoginFunc {
	return func(form xjwt.LoginInForm) (jwt.MapClaims, error) {
		return jwt.MapClaims{"username": form.Username, "password": form.Password}, nil
	}
}

func NewJWTMiddleware(factory xjwt.JWTFactory) *middleware.JWTMiddleware {
	m := middleware.NewJWTMiddleware(middleware.WithFactory(factory))
	return &m
}

var WireSet = wire.NewSet(NewEntry, NewLoginFunc, NewJWTMiddleware, NewJWTFactory, NewRouter, v1.NewUserController, v1.NewLotteryController, v1.NewDefaultController, v1.NewJwtController)
