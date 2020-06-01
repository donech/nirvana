package gin

import (
	"github.com/donech/nirvana/internal/iface/gin/controller"
	"github.com/donech/nirvana/internal/iface/gin/router"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(NewEntry, router.NewRouter, controller.NewUserController)
