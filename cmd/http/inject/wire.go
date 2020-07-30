//+build wireinject

package inject

import (
	"github.com/donech/nirvana/internal/config"
	"github.com/donech/nirvana/internal/conn"
	"github.com/donech/nirvana/internal/domain"
	gin2 "github.com/donech/nirvana/internal/entry/gin"
	"github.com/donech/tool/entry/gin"
	"github.com/donech/tool/xlog"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func providerLogger(conf *config.Config) (logger *zap.Logger, err error) {
	return xlog.New(conf.Log)
}

func InitApplication() (entry *gin.Entry, cleanup func(), err error) {
	wire.Build(
		config.New,
		viper.GetViper,
		gin2.WireSet,
		domain.WireSet,
		providerLogger,
		conn.NewNirvanaDB,
	)
	return &gin.Entry{}, nil, nil
}
