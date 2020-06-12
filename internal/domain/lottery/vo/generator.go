package vo

import (
	"context"

	"github.com/donech/nirvana/internal/domain/lottery/entity"
)

type Generator interface {
	Generate(ctx context.Context, period string) entity.LotteryRecord
}

var GeneratorMap = make(map[TicketType]Generator)

func GetGenerator(ticketType TicketType) Generator {
	out, ok := GeneratorMap[ticketType]
	if !ok {
		return GeneratorMap[TwoToneSphere]
	}
	return out
}
