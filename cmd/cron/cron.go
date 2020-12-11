package cron

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/unknwon/com"

	"github.com/donech/tool/xlog"

	"github.com/donech/nirvana/cmd/cron/inject"
	"github.com/spf13/cobra"
)

var SuperLottoPeriodStart int
var TwoToneSpherePeriodStart int
var TwoToneSphereNumbers string
var SuperLottoNumbers string
var PeriodFile string

func init() {
	Command.PersistentFlags().IntVar(&TwoToneSpherePeriodStart, "tp", 123, "use --tp to set TwoToneSpherePeriodStart")
	Command.PersistentFlags().IntVar(&SuperLottoPeriodStart, "sp", 123, "use --sp to set SuperLottoPeriodStart")
	Command.PersistentFlags().StringVar(&TwoToneSphereNumbers, "tn", "", "use --tn to set TwoToneSphereNumbers")
	Command.PersistentFlags().StringVar(&SuperLottoNumbers, "sn", "", "use --sn to set SuperLottoNumbers")
	Command.PersistentFlags().StringVar(&PeriodFile, "pf", "", "use --pf to set PeriodFile")
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
		checkoutPeriodFile()
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
			logPeriodFile(app.Tp, app.Sp)
			log.Println("end task", "下一期为:", app.Tp, app.Sp)
		}
		task()
		return nil
	},
}

func checkoutPeriodFile() {
	content, err := ioutil.ReadFile(PeriodFile)
	if err != nil {
		log.Printf("read file %s error %v", PeriodFile, err)
		return
	}
	period := strings.Split(string(content), "|")
	if len(period) == 2 {
		log.Println("存在历史记录忽略期数配置", "记录文件为:", PeriodFile)
		TwoToneSpherePeriodStart, SuperLottoPeriodStart = com.StrTo(period[0]).MustInt(), com.StrTo(period[1]).MustInt()
	}
}

func logPeriodFile(tPeriod, sPeriod int) {
	t := "%d|%d"
	content := fmt.Sprintf(t, tPeriod, sPeriod)
	f, err := os.OpenFile(PeriodFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err.Error())
	}

	_, err = f.Write([]byte(content))
	if err != nil {
		log.Println(err.Error())
	}
	f.Close()
}
