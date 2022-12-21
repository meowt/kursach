package config

import (
	"github.com/spf13/viper"
	"log"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./pkg/config")
	viper.AddConfigPath("./pkg/database")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	log.Println("Config initialised")
	return
}
