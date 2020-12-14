//+build wireinject

package inject

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/donech/nirvana/internal/domain/lottery/vo"

	"github.com/donech/nirvana/internal/config"
	"github.com/donech/nirvana/internal/conn"
	"github.com/donech/tool/xlog"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var SuperLottoPeriodTep = "%d-%d"
var TwoToneSpherePeriodTep = "%d%d"

type Entry struct {
	conf   *config.Config
	logger *zap.Logger
	client *conn.LotteryClient
	Tp     int
	Sp     int
	Tn     string
	Sn     string
}

func (e *Entry) Run() error {
	weekday := time.Now().Weekday()
	switch weekday {
	case time.Monday, time.Wednesday, time.Friday:
		e.checkTwoToneSphere()
	case time.Tuesday, time.Thursday, time.Sunday:
		e.checkSupperLotto()
	default:
		log.Printf("Saturday don't need to search result")
	}
	return nil
}

func (e *Entry) checkTwoToneSphere() error {
	period := fmt.Sprintf(TwoToneSpherePeriodTep, time.Now().Year(), e.Tp)
	if e.Tn == "" {
		log.Printf("无 双色球号码 返回")
		return nil
	}
	number := e.client.GetTwoToneSphere(context.Background(), period)
	calculator, ok := vo.CalculatorMap[vo.TwoToneSphere]
	if !ok {
		log.Printf("无 双色球计算器 返回")
		return nil
	}
	level, price := calculator.Calculate(e.Tn, number)
	log.Printf("双色球%s期开奖结果为:%s,你的号码为:%s, 中奖情况 level:%d, price:%d", period, number, e.Tn, level, price)
	e.Tp++
	return nil
}

func (e *Entry) checkSupperLotto() error {
	period := fmt.Sprintf(SuperLottoPeriodTep, time.Now().Year(), e.Sp)
	if e.Sn == "" {
		log.Printf("无 大乐透号码 返回")
		return nil
	}
	number := e.client.GetSupperLotto(context.Background(), period)
	calculator, ok := vo.CalculatorMap[vo.SuperLotto]
	if !ok {
		log.Printf("无 大乐透计算器 返回")
		return nil
	}
	level, price := calculator.Calculate(e.Sn, number)
	log.Printf("大乐透%s期开奖结果为:%s,你的号码为:%s, 中奖情况 level:%d, price:%d", period, number, e.Sn, level, price)
	e.Sp++
	return nil
}

func (e *Entry) Stop(ctx context.Context) error {
	panic("implement me")
}

func NewEntry(conf *config.Config, logger *zap.Logger, client *conn.LotteryClient) *Entry {
	return &Entry{
		conf:   conf,
		logger: logger,
		client: client,
	}
}

func providerLogger(conf *config.Config) (logger *zap.Logger, err error) {
	return xlog.New(conf.Log)
}

func InitApplication() (entry *Entry, cleanup func(), err error) {
	wire.Build(
		config.New,
		viper.GetViper,
		providerLogger,
		conn.NewLotteryClient,
		NewEntry,
	)
	return &Entry{}, nil, nil
}
