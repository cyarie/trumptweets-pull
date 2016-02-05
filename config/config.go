package config

import (
	"github.com/spf13/viper"
	"fmt"
)

type Conf struct {
	Development    AppConfig    `json:"development"`
}

type AppConfig struct {
	AppName    string         `json:"appName"`
	Twitter    TwitterConf    `json:"twitter"`
	Db         DbConf         `json:"database"`
}

type TwitterConf struct {
	ConsumerKey    string    `json:"consumerKey"`
	ConsumerSecret string    `json:"consumerSecret"`
	Token          string    `json:"token"`
	TokenSecret    string    `json:"tokenSecret"`
}

type DbConf struct {
	Host           string    `json:"host"`
	Name           string    `json:"name"`
	User           string    `json:"user"`
	Password       string    `json:"password"`
}

var Config Conf

func LoadConfig() {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$GOPATH/src/github.com/cyarie/trump_tweets/config/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err.Error())
	}
}
