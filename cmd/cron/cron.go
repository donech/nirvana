package cron

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/coreos/etcd/pkg/fileutil"

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
var LockFileTemp string = "/tmp/period-%s"

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
		lockfile := fmt.Sprintf(LockFileTemp, time.Now().Format("2006-01-02"))
		if checkLockFile(lockfile) {
			log.Printf("It's already runed this day!")
			return nil
		}
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
			logPeriodFile(app.Tp, app.Sp)
			log.Printf("task end  下一期双色球为: %d, 大乐透: %d", app.Tp, app.Sp)
		}
		task()
		createLockFile(lockfile)
		return nil
	},
}

// checkLockFile 检查当天是否执行过，执行过不再执行
func checkLockFile(filename string) bool {
	return fileutil.Exist(filename)
}

func createLockFile(filename string) {
	_, err := os.Create(filename)
	if err != nil {
		log.Printf("生成 LockFile 失败: %s", err)
	}

}

func checkoutPeriodFile() {
	if !fileutil.Exist(PeriodFile) {
		log.Printf("无历史查询文件: %s,  使用命令行参数", PeriodFile)
		return
	}
	content, err := ioutil.ReadFile(PeriodFile)
	if err != nil {
		log.Printf("read file %s error %v", PeriodFile, err)
		return
	}
	if content[len(content)-1] == '\n' {
		content = content[:len(content)-1]
	}
	period := strings.Split(string(content), "|")
	if len(period) == 2 {
		log.Println("存在历史记录忽略期数配置", "记录文件为:", PeriodFile)
		log.Printf("原始 period %#v", period)
		TwoToneSpherePeriodStart, SuperLottoPeriodStart = com.StrTo(period[0]).MustInt(), com.StrTo(period[1]).MustInt()
		log.Printf("双色球期数: %d, 大乐透期数 %d", TwoToneSpherePeriodStart, SuperLottoPeriodStart)
	}
}

func logPeriodFile(tPeriod, sPeriod int) {
	t := "%d|%d"
	content := fmt.Sprintf(t, tPeriod, sPeriod)
	f, err := os.OpenFile(PeriodFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println(err.Error())
	}

	_, err = f.Write([]byte(content))
	if err != nil {
		log.Println(err.Error())
	}
	f.Close()
}
