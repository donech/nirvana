package gin

import (
	v1 "github.com/donech/nirvana/internal/entry/gin/api/v1"
	"github.com/donech/nirvana/internal/entry/gin/router"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(NewEntry, router.NewRouter, v1.NewUserController, v1.NewLotteryController)
