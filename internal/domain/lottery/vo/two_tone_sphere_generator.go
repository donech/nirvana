package vo

import (
	"context"

	"github.com/donech/nirvana/internal/conn"
	"github.com/donech/nirvana/internal/domain/lottery/entity"
)

type TwoToneSphereGenerator struct {
}

func (t TwoToneSphereGenerator) Generate(ctx context.Context, period string) entity.LotteryRecord {
	client := conn.LotteryClient{}
	ssq := client.GetTwoToneSphere(ctx, period)
	return entity.LotteryRecord{
		Number: ssq,
		Period: period,
		Type:   string(TwoToneSphere),
	}
}

func init() {
	GeneratorMap[TwoToneSphere] = TwoToneSphereGenerator{}
}
