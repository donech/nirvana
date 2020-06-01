package router

import (
	"github.com/donech/nirvana/internal/iface/gin/controller"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	RegisterRoute(engine *gin.Engine)
}

func NewRouter(userController *controller.UserController) *Router {
	controllers := []Controller{userController}
	return &Router{controllers: controllers}
}

type Router struct {
	controllers []Controller
}

func (r Router) Init(engine *gin.Engine) {
	for _, c := range r.controllers {
		c.RegisterRoute(engine)
	}
}
