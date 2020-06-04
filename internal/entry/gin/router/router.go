package router

import (
	v1 "github.com/donech/nirvana/internal/entry/gin/api/v1"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	RegisterRoute(engine *gin.RouterGroup)
}

func NewRouter(userController *v1.UserController) *Router {
	controllers := []Controller{userController}
	return &Router{controllers: controllers}
}

type Router struct {
	controllers []Controller
}

func (r Router) Init(engine *gin.Engine) {
	rootGroup := engine.Group("/")
	for _, c := range r.controllers {
		c.RegisterRoute(rootGroup)
	}
}
