package vo

import (
	"context"

	"github.com/donech/nirvana/internal/conn"

	"github.com/donech/nirvana/internal/domain/lottery/entity"
)

func init() {
	GeneratorMap[SuperLotto] = SuperLottoGenerator{}
}

type SuperLottoGenerator struct{}

func (SuperLottoGenerator) Generate(ctx context.Context, period string) entity.LotteryRecord {
	client := conn.LotteryClient{}
	supper := client.GetSupperLotto(ctx, period)
	return entity.LotteryRecord{
		Number: supper,
		Period: period,
		Type:   string(SuperLotto),
	}
}
