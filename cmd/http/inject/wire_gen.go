// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package inject

import (
	"github.com/donech/nirvana/internal/config"
	"github.com/donech/nirvana/internal/conn"
	"github.com/donech/nirvana/internal/domain/user/service"
	"github.com/donech/nirvana/internal/entry/gin"
	"github.com/donech/nirvana/internal/entry/gin/api/v1"
	"github.com/donech/nirvana/internal/entry/gin/router"
	"github.com/donech/tool/xlog"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func InitApplication() (*gin.Entry, func(), error) {
	viperViper := viper.GetViper()
	configConfig := config.New(viperViper)
	nirvanaDB, cleanup := conn.NewNirvanaDB(configConfig)
	simpleService := service.NewSimpleService(nirvanaDB)
	userController := v1.NewUserController(simpleService)
	routerRouter := router.NewRouter(userController)
	logger, err := providerLogger(configConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	entry := gin.NewEntry(configConfig, routerRouter, logger)
	return entry, func() {
		cleanup()
	}, nil
}

// wire.go:

func providerLogger(conf *config.Config) (logger *zap.Logger, err error) {
	return xlog.New(conf.Log)
}
