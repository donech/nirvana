package conn

import (
	"context"
	"log"
	"sync"
	"testing"

	"github.com/donech/tool/xlog"

	"github.com/stretchr/testify/assert"
)

func init() {
	conf := xlog.Config{
		ServiceName: "xlog-test",
		Level:       "info",
		LevelColor:  true,
		Format:      "json",
		Stdout:      true,
		File: xlog.FileLogConfig{
			Filename:   "test.log",
			LogRotate:  true,
			MaxSize:    20,
			MaxAge:     20,
			MaxBackups: 10,
			BufSize:    20,
		},
		EncodeKey: xlog.EncodeKeyConfig{},
		SentryDSN: "",
	}
	_, err := xlog.New(conf)
	if err != nil {
		log.Fatal("创建 ginzap.logger 失败")
	}
}

func TestLotteryClient_GetTwoToneSphere(t *testing.T) {
	client := NewLotteryClient()
	tests := []struct {
		desc     string
		period   string
		expected string
	}{
		{desc: "2020048期", period: "2020048", expected: "12,14,18,23,30,32|02"},
		{desc: "2020047期", period: "2020047", expected: "04,10,17,19,28,32|01"},
	}
	wg := sync.WaitGroup{}
	for _, tt := range tests {
		wg.Add(1)
		go func(tt struct {
			desc     string
			period   string
			expected string
		}) {
			result := client.GetTwoToneSphere(context.Background(), tt.period)
			assert.Equal(t, tt.expected, result, "%s: unexpected :%s expect:%s", tt.desc, result, tt.expected)
			wg.Done()
		}(tt)
	}
	wg.Wait()
}

func TestLotteryClient_GetSupperLotto(t *testing.T) {
	tests := []struct {
		name   string
		period string
		want   string
	}{
		{name: "2020-123期", period: "2020-123", want: "01,07,11,15,21|04,11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LotteryClient{}
			if got := c.GetSupperLotto(context.Background(), tt.period); got != tt.want {
				t.Errorf("GetSupperLotto() = %v, want %v", got, tt.want)
			}
		})
	}
}
