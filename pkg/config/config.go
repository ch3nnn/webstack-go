package config

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func NewConfig(p string) *viper.Viper {
	conf := viper.New()
	conf.AutomaticEnv()

	envConf := conf.GetString("APP_CONF")
	if envConf != "" {
		p = envConf
	}

	if _, err := os.Stat(p); os.IsNotExist(err) {
		panic(errors.Errorf("config file not found: %s", p))
	}

	conf.SetConfigFile(p)

	if err := conf.ReadInConfig(); err != nil {
		panic(errors.Errorf("failed to read config file: %s", err))
	}

	return conf
}
