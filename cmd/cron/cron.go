package cron

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"

	"github.com/donech/tool/xlog"

	"github.com/donech/nirvana/cmd/cron/inject"
	"github.com/spf13/cobra"
)

var SuperLottoPeriodStart int
var TwoToneSpherePeriodStart int
var TwoToneSphereNumbers string
var SuperLottoNumbers string

func init() {
	Command.PersistentFlags().IntVar(&TwoToneSpherePeriodStart, "tp", 123, "use -tp to set TwoToneSpherePeriodStart")
	Command.PersistentFlags().IntVar(&SuperLottoPeriodStart, "sp", 123, "use -sp to set SuperLottoPeriodStart")
	Command.PersistentFlags().StringVar(&TwoToneSphereNumbers, "tn", "", "use -tn to set TwoToneSphereNumbers")
	Command.PersistentFlags().StringVar(&SuperLottoNumbers, "sn", "", "use -sn to set SuperLottoNumbers")
}

var Command = &cobra.Command{
	Use:   "cron",
	Short: "Timing request and sending system notification\n",
	Long:  `Timing request and sending system notification`,
	RunE: func(cmd *cobra.Command, args []string) error {
		app, clean, err := inject.InitApplication()
		if err != nil {
			log.Fatalf("init application error %v", err)
		}
		defer clean()
		app.Tp = TwoToneSpherePeriodStart
		app.Sp = SuperLottoPeriodStart
		app.Tn = TwoToneSphereNumbers
		app.Sn = SuperLottoNumbers
		xlog.S(context.Background()).Infof("app is %#v", app)
		task := func() {
			log.Println("begin task")
			err = app.Run()
			if err != nil {
				log.Fatalf("run application error %v", err)
			}
			app.Tp++
			app.Sp++
			log.Println("end task")
		}
		crontab := cron.New()
		_, err = crontab.AddFunc("* * * * *", task)
		if err != nil {
			log.Fatalf("add crontab error, %#v", err)
		}
		crontab.Start()
		signalCh := make(chan os.Signal, 2)
		signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

		select {
		case <-signalCh:
			log.Println("Interrupt")
		}
		return nil
	},
}
