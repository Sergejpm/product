package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPPort             uint   `envconfig:"HTTP_PORT" default:"8080"`
	LogLevel             string `envconfig:"LOG_LEVEL" default:"DEBUG"`
	DBConnectionString   string `envconfig:"CONNECTION_STRING"`
	DBMaxIdleConnections int    `envconfig:"DB_MAX_IDLE_CONNECTIONS" default:"1"`
	DBMaxOpenConnections int    `envconfig:"DB_MAX_OPEN_CONNECTIONS" default:"1"`
	TokenKey             []byte `envconfig:"TOKEN_KEY" default:"example"`
}

func Load() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
