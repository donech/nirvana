package repository

import (
	"context"

	"github.com/donech/nirvana/internal/conn"
	"github.com/donech/nirvana/internal/domain/lottery/entity"
	"github.com/donech/tool/xdb"
)

func NewRecordRepository(db conn.NirvanaDB) *RecordRepository {
	return &RecordRepository{Repository: xdb.Repository{DB: db}}
}

type RecordRepository struct {
	xdb.Repository
}

func (r *RecordRepository) ItemByPeriodAndType(ctx context.Context, period, tp string) (record entity.LotteryRecord, err error) {
	err = xdb.Trace(ctx, r.DB).Where("period = ? and type = ?", period, tp).First(&record).Error
	return record, err
}
