package service

import (
	"context"
	"errors"

	"github.com/donech/nirvana/internal/domain/lottery/vo"

	"github.com/jinzhu/gorm"

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

func (s *LotteryService) TicketCheck(ctx context.Context, id int64) (entity.LotteryTicket, error) {
	ticket, err := s.TicketByID(ctx, id)
	if err != nil || ticket.IsChecked() {
		return ticket, err
	}
	record, err := s.RecordByPeriodAndType(ctx, ticket.Period, ticket.TicketType)
	if err != nil {
		return ticket, err
	}
	calculator := vo.GetCalculator(vo.TicketType(record.Type))
	level, price := calculator.Calculate(ticket.Number, record.Number)
	ticket.Level = level
	ticket.Price = price
	ticket.Status = entity.TicketCheckedStatus
	err = s.ticketRepository.Save(ctx, &ticket)
	if err != nil {
		return ticket, err
	}
	return ticket, err
}

func (s *LotteryService) CreateTicket(ctx context.Context, ticket *entity.LotteryTicket) error {
	return s.ticketRepository.Create(ctx, ticket)
}

func (s *LotteryService) CreateRecord(ctx context.Context, record *entity.LotteryRecord) error {
	return s.ticketRepository.Create(ctx, record)
}

func (s *LotteryService) RecordByPeriodAndType(ctx context.Context, period, tp string) (entity.LotteryRecord, error) {
	e, err := s.recordRepository.ItemByPeriodAndType(ctx, period, tp)
	return e, err
}

func (s *LotteryService) GenerateRecordByPeriodAndType(ctx context.Context, period, tp string) (entity.LotteryRecord, error) {
	e, err := s.RecordByPeriodAndType(ctx, period, tp)
	if err != nil && err != gorm.ErrRecordNotFound {
		return entity.LotteryRecord{}, err
	}
	if err == gorm.ErrRecordNotFound {
		e = vo.GetGenerator(vo.TicketType(tp)).Generate(ctx, period)
		if e.Number == "" {
			return e, errors.New("未查到彩票信息")
		}
		err = s.CreateRecord(ctx, &e)
	}
	return e, err
}

func (s LotteryService) Migration() {
	s.ticketRepository.DB.AutoMigrate(entity.LotteryTicket{})
	s.recordRepository.DB.AutoMigrate(entity.LotteryRecord{})
}
