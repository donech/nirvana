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
	if ssq.State != 0 {
		return entity.LotteryRecord{}
	}
	return entity.LotteryRecord{
		Number: ssq.Result[0].Red + "|" + ssq.Result[0].Blue,
		Period: period,
		Type:   string(TwoToneSphere),
	}
}

func init() {
	GeneratorMap[TwoToneSphere] = TwoToneSphereGenerator{}
}
