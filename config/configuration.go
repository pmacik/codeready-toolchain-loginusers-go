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

// Config initiate conifiguration, set default values and loads configuration from default file.
func Config() Configuration {
	return NewConfig("config", ".", "yml")
}

// NewConfig initiate conifiguration, set default values and loads configuration from file.
func NewConfig(path string, name string, configType string) Configuration {
	v := viper.New()

	v.SetConfigName(name)
	v.AddConfigPath(path)
	v.AutomaticEnv()
	v.SetConfigType(configType)
	err := v.ReadInConfig()
	common.CheckErr(err)
	var cfg Configuration
	SetDefaults(&cfg)
	err = v.Unmarshal(&cfg)
	common.CheckErr(err)
	return cfg
}
