package config

import (
	"github.com/pmacik/loginusers-go/common"
	"github.com/spf13/viper"
)

// Configuration for loginusers
type Configuration struct {
	Auth         AuthConfiguration
	Chromedriver ChromedriverConfiguration
	Users        UsersConfiguration
}

// DefaultConfig creates a default configuration set
func DefaultConfig() Configuration {
	var cfg Configuration
	SetDefaults(&cfg)
	return cfg
}

// Config initiate conifiguration, set default values and loads configuration from file.
func Config() Configuration {
	v := viper.New()

	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	common.CheckErr(err)
	var cfg Configuration
	SetDefaults(&cfg)
	err = v.Unmarshal(&cfg)
	return cfg
}
