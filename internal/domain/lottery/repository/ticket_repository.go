package repository

import (
	"context"

	"github.com/donech/nirvana/internal/conn"
	"github.com/donech/nirvana/internal/domain/lottery/entity"
	"github.com/donech/tool/xdb"
)

func NewTicketRepository(db conn.NirvanaDB) *TicketRepository {
	return &TicketRepository{Repository: xdb.Repository{DB: db}}
}

type TicketRepository struct {
	xdb.Repository
}

func (r *TicketRepository) ItemsByPeriod(ctx context.Context, period string) (entities []entity.LotteryTicket, err error) {
	err = xdb.Trace(ctx, r.DB).Where("period = ?", period).Find(&entities).Error
	return entities, err
}

func (r *TicketRepository) ItemByID(ctx context.Context, id int64) (entity entity.LotteryTicket, err error) {
	err = xdb.Trace(ctx, r.DB).Where("id = ?", id).Find(&entity).Error
	return entity, err
}
