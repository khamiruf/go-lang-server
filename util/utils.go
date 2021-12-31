package util

import (
	"github.com/spf13/viper"
	"go-lang-server/config"
)

func LoadConfig(path string) (*config.GeneralConfig, error) {
	var config *config.GeneralConfig

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)

	return config, nil
}
