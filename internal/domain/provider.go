package domain

import (
	"github.com/donech/nirvana/internal/domain/user/service"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(service.NewSimpleService)
