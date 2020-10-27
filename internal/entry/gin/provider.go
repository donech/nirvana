package gin

import (
	"context"
	"log"

	"github.com/donech/nirvana/internal/entry/gin/api"

	xlog "github.com/donech/tool/xlog"

	"github.com/donech/nirvana/internal/domain/user/service"

	"github.com/dgrijalva/jwt-go"

	"github.com/donech/tool/entry/gin/middleware"

	"github.com/donech/nirvana/internal/config"
	_ "github.com/donech/nirvana/internal/entry/gin/docs"
	"github.com/donech/tool/entry/gin"
	"github.com/donech/tool/xjwt"

	"github.com/google/wire"
	"go.uber.org/zap"
)

func NewRouter(
	u *api.UserController,
	l *api.LotteryController,
	d *api.DefaultController,
	jwt *api.JwtController,
) gin.Router {
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

func NewLoginFunc(s *service.UserService) xjwt.LoginFunc {
	return func(ctx context.Context, form xjwt.LoginForm) (jwt.MapClaims, error) {
		user, err := s.Login(ctx, form.Username, form.Password)
		if err != nil {
			xlog.S(ctx).Info("登录失败: LoginForm=%#v, err=%#v", form, err)
			return nil, err
		}
		return jwt.MapClaims{
			"id":       user.ID,
			"username": user.Name,
			"email":    user.Email,
			"phone":    user.Phone,
			"status":   user.Status,
		}, nil
	}
}

func NewJWTMiddleware(factory xjwt.JWTFactory) *middleware.JWTMiddleware {
	m := middleware.NewJWTMiddleware(middleware.WithFactory(factory))
	return &m
}

var WireSet = wire.NewSet(NewEntry, NewLoginFunc, NewJWTMiddleware, NewJWTFactory, NewRouter, api.NewUserController, api.NewLotteryController, api.NewDefaultController, api.NewJwtController)
