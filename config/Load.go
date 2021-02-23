package config

import "github.com/spf13/viper"

func Load() (*viper.Viper, error) {
	conf := viper.GetViper()
	conf.AddConfigPath(".")
	conf.SetConfigName("conf")

	err := conf.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return conf, nil
}