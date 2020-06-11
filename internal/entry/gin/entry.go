package gin

import (
	"context"
	"log"
	"net/http"
	"time"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/donech/tool/xlog/ginzap"

	"go.uber.org/zap"

	"github.com/donech/nirvana/internal/entry/gin/router"

	"github.com/donech/nirvana/internal/config"
	_ "github.com/donech/nirvana/internal/entry/gin/docs"

	"github.com/gin-gonic/gin"
)

var E *Entry

func NewEntry(conf *config.Config, router *router.Router, logger *zap.Logger) *Entry {
	engine := gin.New()
	engine.Use(ginzap.GinZap(zap.L(), time.RFC3339, true, conf.Application.Mod))
	engine.Use(ginzap.RecoveryWithZap(zap.L(), true))
	engine.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "health", "connection_num": ginzap.GetConnectionNum(),
		})
	})
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	E = &Entry{
		conf:   conf,
		engine: engine,
		router: router,
		logger: logger,
	}
	return E
}

type Entry struct {
	conf   *config.Config
	engine *gin.Engine
	router *router.Router
	logger *zap.Logger
	srv    *http.Server
}

func (e *Entry) Run() error {
	srv := &http.Server{
		Addr:    e.conf.Gin.Addr,
		Handler: e.engine,
	}
	e.router.Init(e.engine)
	go func() {
		log.Println("start http server at ", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	e.srv = srv
	return nil
}

func (e *Entry) Stop(ctx context.Context) error {
	err := e.srv.Shutdown(ctx)
	if err != nil {
		log.Println("Http Server Shutdown failed: ", err)
		return err
	}
	log.Println("Http Server exiting")
	return nil
}
