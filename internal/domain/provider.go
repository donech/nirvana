package domain

import (
	repository2 "github.com/donech/nirvana/internal/domain/lottery/repository"
	"github.com/donech/nirvana/internal/domain/lottery/service"
	"github.com/donech/nirvana/internal/domain/user/repository"
	service2 "github.com/donech/nirvana/internal/domain/user/service"
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewLotteryService,
	repository2.NewTicketRepository,
	repository2.NewRecordRepository,
	service2.NewUserService,
)
