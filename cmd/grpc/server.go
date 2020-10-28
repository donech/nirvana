//Package grpc http server command
/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package grpc

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/donech/nirvana/cmd/grpc/inject"
	"github.com/donech/nirvana/internal/common"
	"github.com/donech/nirvana/internal/config"
	"github.com/donech/tool/entry"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ServerCmd represents the server command
var ServerCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Server grpc",
	Long:  `Server grpc`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

func run() error {
	en, clean, err := inject.InitApplication()
	common.SetKey(config.C.Application.Key)
	defer func() {
		if clean != nil {
			clean()
		}
	}()

	if err != nil {
		log.Println("init application failed :", err.Error())
		return err
	}

	err = en.Run()
	if err != nil {
		log.Println("en run failed :", err.Error())
	}
	handleSignal(en)
	return err
}

func handleSignal(en entry.Entry) {
	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
WAIT:
	s := <-signalCh
	if s == syscall.SIGHUP {
		config.New(viper.GetViper())
		log.Println("reload config")
		goto WAIT
	}
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := en.Stop(ctx); err != nil {
		log.Fatal("Entry stop err :", err)
	}
	log.Println("Entry exiting")
}
