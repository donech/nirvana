//+build wireinject

package inject

import (
	"github.com/donech/nirvana/internal/config"
	grpc2 "github.com/donech/nirvana/internal/entry/grpc"
	"github.com/donech/tool/entry/grpc"
	"github.com/donech/tool/xlog"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func providerLogger(conf *config.Config) (logger *zap.Logger, err error) {
	return xlog.New(conf.Log)
}

func InitApplication() (entry *grpc.Entry, cleanup func(), err error) {
	wire.Build(
		config.New,
		viper.GetViper,
		grpc2.NewRegisteServer,
		providerLogger,
		grpc2.NewEntry,
	)
	return &grpc.Entry{}, nil, nil
}
