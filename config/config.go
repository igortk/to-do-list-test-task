package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/pkg/errors"
)

type Config struct {
	HttpConfig     HttpConfig     `envPrefix:"HTTP_"`
	PostgresConfig PostgresConfig `envPrefix:"PG_"`
}

type HttpConfig struct {
	Port string `env:"PORT" envDefault:":8888"`
}

type PostgresConfig struct {
	Dsn string `env:"DSN" envDefault:"host=localhost user=postgres password=password dbname=to-do-list port=5432 sslmode=disable"`
}

func Read() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.Wrap(err, "error parsing config from env")
	}

	return &cfg, nil
}
