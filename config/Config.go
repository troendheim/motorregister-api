package config

import (
	"github.com/spf13/viper"
	"fmt"
)

func PrepareConfig() {

	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("app")

	viper.SetDefault("port", 8999)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
