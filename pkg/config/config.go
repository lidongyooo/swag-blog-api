package config

import (
	"github.com/spf13/viper"
)

var Viper *viper.Viper

func init()  {
	Viper = viper.New()

	Viper.SetConfigName(".env")
	Viper.SetConfigType("env")
	Viper.AddConfigPath(".")

	err := Viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(err.Error())
	}

	Viper.SetEnvPrefix("swagblog")
	Viper.AutomaticEnv()
}



