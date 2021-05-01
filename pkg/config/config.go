package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}
