package config

import (
	"it-test/config"

	"github.com/spf13/viper"
)

type ViperConfig struct {
}

func (c *ViperConfig) LoadConfig(path string, cfg config.Config) error {
	for k, v := range cfg.Defaults() {
		viper.SetDefault(k, v)
	}
	viper.AddConfigPath(path)
	viper.SetConfigName(cfg.Name())
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&cfg)
}
