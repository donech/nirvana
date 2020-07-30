package gin

import (
	"github.com/donech/nirvana/internal/config"
	v1 "github.com/donech/nirvana/internal/entry/gin/api/v1"
	_ "github.com/donech/nirvana/internal/entry/gin/docs"
	"github.com/donech/tool/entry/gin"

	"github.com/google/wire"
	"go.uber.org/zap"
)

func NewRouter(u *v1.UserController, l *v1.LotteryController, d *v1.DefaultController) gin.Router {
	r := &gin.DefaultRouter{}
	r.RegisterController(u)
	r.RegisterController(l)
	r.RegisterController(d)
	return r
}

func NewEntry(conf *config.Config, router gin.Router, logger *zap.Logger) *gin.Entry {
	confG := &gin.Config{
		Mod:  conf.Application.Mod,
		Addr: conf.Gin.Addr,
	}
	return gin.NewEntry(confG, router, logger)
}

var WireSet = wire.NewSet(NewEntry, NewRouter, v1.NewUserController, v1.NewLotteryController, v1.NewDefaultController)
