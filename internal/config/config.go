package config

import (
	"log"

	"github.com/donech/tool/xjwt"

	"github.com/donech/tool/xdb"

	"github.com/donech/tool/xlog"

	"github.com/spf13/viper"
)

var C *Config

type Config struct {
	Application ApplicationConfig `yaml:"application"`
	Gin         GinConfig         `yaml:"gin"`
	Log         xlog.Config       `yaml:"log"`
	NirvanaDB   xdb.Config        `yaml:"nirvanaDB"`
	JWT         xjwt.Config       `yaml:"jwt"`
}

type ApplicationConfig struct {
	Name string `yaml:"name"`
	Mod  string `yaml:"mod"`
	Key  string `yaml:"key"`
}

type GinConfig struct {
	Addr string `yaml:"addr"`
}

func New(viper *viper.Viper) *Config {
	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatalln("can't unmarshal viper to Config :", err)
	}
	return C
}
