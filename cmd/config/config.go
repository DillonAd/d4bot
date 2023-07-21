package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	OtelEndpoint string
	OtelInsecure bool `default:"false"`
	DiscordToken string
}

func NewConfig() (Config, error) {
	var config Config
	err := envconfig.Process("d4bot", &config)

	if err != nil {
		return config, fmt.Errorf("error fetching env vars: %v", err)
	}

	return config, nil
}
