//Package cmd comands
/*
Copyright © 2020 5412 <solarpwx@yeah.net>

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
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/donech/nirvana/cmd/grpc"
	"github.com/donech/nirvana/cmd/http"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nirvana",
	Short: "A simple web service provider",
	Long:  `A simple web service provider combines gin framwork and Grpc service`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Welcome to use nirvana, see more information whit -h flag")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "cmd/app.yaml", "config file (default is cmd/app.yaml)")
	rootCmd.AddCommand(http.ServerCmd)
	rootCmd.AddCommand(grpc.ServerCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	log.Println("Using config file:", cfgFile)
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Panicln("Read config file failed:", viper.ConfigFileUsed(), err.Error())
	}
}

//RegistCommand Regist a Command to rootCommand
func RegistCommand(c *cobra.Command) {
	rootCmd.AddCommand(c)
}
