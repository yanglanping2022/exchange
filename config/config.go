package config

import (
	"github.com/BurntSushi/toml"
)

type binanceInfo struct {
	APIKey    string `toml:"apikey"`
	SecretKey string `toml:"secretkey"`
	BaseURL   string `toml:"baseurl"`
}

type gateioInfo struct {
	APIKey    string `toml:"apikey"`
	SecretKey string `toml:"secretkey"`
}

type Config struct {
	Binance binanceInfo `toml:"binance"`
	Gateio  gateioInfo  `toml:"gateio"`
}

var Conf Config

func InitConf() {
	if _, err := toml.DecodeFile("./config.toml", &Conf); err != nil {
		panic(err)
	}
}
