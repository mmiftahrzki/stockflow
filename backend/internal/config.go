package internal

import "github.com/spf13/viper"

type AppConfig struct {
	Port       uint16
}

type baseConfig struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
}

type DatabaseConfig struct {
	User          string
	Password      string
	Host          string
	Port          uint16
	Name          string
	MaxConnection int
}

func LoadConfig() (baseConfig, error) {
	var cfg baseConfig

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./../..")

	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
