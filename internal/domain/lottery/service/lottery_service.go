package service

import (
	"context"

	"github.com/donech/nirvana/internal/domain/lottery/entity"
	"github.com/donech/nirvana/internal/domain/lottery/repository"
)

func NewLotteryService(ticketRepository *repository.TicketRepository, recordRepository *repository.RecordRepository) *LotteryService {
	return &LotteryService{ticketRepository: ticketRepository, recordRepository: recordRepository}
}

type LotteryService struct {
	ticketRepository *repository.TicketRepository
	recordRepository *repository.RecordRepository
}

func (s *LotteryService) TicketsByPeriod(ctx context.Context, period string) ([]entity.LotteryTicket, error) {
	return s.ticketRepository.ItemsByPeriod(ctx, period)
}

func (s *LotteryService) TicketByID(ctx context.Context, id int64) (entity.LotteryTicket, error) {
	return s.ticketRepository.ItemByID(ctx, id)
}

func (s *LotteryService) CreateTicket(ctx context.Context, ticket *entity.LotteryTicket) error {
	return s.ticketRepository.Create(ctx, ticket)
}

func (s *LotteryService) CreateRecord(ctx context.Context, record entity.LotteryRecord) error {
	return s.ticketRepository.Create(ctx, record)
}

func (s *LotteryService) RecordByPeriodAndType(ctx context.Context, period, tp string) (entity.LotteryRecord, error) {
	return s.recordRepository.ItemByPeriodAndType(ctx, period, tp)
}

func (s LotteryService) Migration() {
	s.ticketRepository.DB.AutoMigrate(entity.LotteryTicket{})
	s.recordRepository.DB.AutoMigrate(entity.LotteryRecord{})

}
