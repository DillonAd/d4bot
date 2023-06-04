package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Tracing     *Tracing     `mapstructure:",squash"`
	Healthcheck *Healthcheck `mapstructure:",squash"`
	Bot         *Bot         `mapstructure:",squash"`
}

type Tracing struct {
	OtelEndpoint string `mapstructure:"OTEL_ENDPOINT"`
}

type Healthcheck struct {
	ApiPath string `mapstructure:"HEALTH_PATH"`
}

type Bot struct {
	Token string `mapstructrue:"BOT_TOKEN"`
}

func Read() (*Config, error) {
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("error reading config file: %v", err)
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %v", err)
	}

	return cfg, nil
}
