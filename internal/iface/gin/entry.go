package gin

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/donech/tool/xlog/ginzap"

	"go.uber.org/zap"

	"github.com/donech/nirvana/internal/iface/gin/router"

	"github.com/donech/nirvana/internal/config"
	"github.com/gin-gonic/gin"
)

var E *Entry

func NewEntry(conf *config.Config, router *router.Router, logger *zap.Logger) *Entry {
	engine := gin.New()
	engine.Use(ginzap.GinZap(zap.L(), time.RFC3339, true, conf.Application.Mod))
	engine.Use(ginzap.RecoveryWithZap(zap.L(), true))
	engine.GET("/redness", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"ping": "pong",
		})
	})
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
}

func (e Entry) Run() error {
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
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	return nil
}
