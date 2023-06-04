package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Tracing     Tracing     `mapstructure:",squash"`
	Healthcheck Healthcheck `mapstructure:",squash"`
	Bot         Bot         `mapstructure:",squash"`
}

type Tracing struct {
	OtelEndpoint string `mapstructure:"OTEL_ENDPOINT"`
}

type Healthcheck struct {
	ApiPath string `mapstructure:"HEALTH_PATH"`
}

type Bot struct {
	Token string `mapstructure:"BOT_TOKEN"`
}

func Read() (*Config, error) {
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %v", err)
	}
	fmt.Printf("returning config:: %+v", config)
	return config, nil
}
