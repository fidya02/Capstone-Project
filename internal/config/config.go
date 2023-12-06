package config

import (
	"errors"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Config is a config
type Config struct {
	Port     string         `env:"PORT" envDefault:"8080"`
	Postgres PostgresConfig `envPrefix:"POSTGRES_"`
	JWT      JwtConfig      `envPrefix:"JWT_"`
	Session  SessionConfig  `envPrefix:"SESSION_"`
}

// JwtConfig is a config for jwt
type JwtConfig struct {
	SecretKey string `env:"SECRET_KEY"`
}

type SessionConfig struct {
	SecretKey string `env:"SECRET_KEY"`
}

// PostgresConfig is a config for postgres
type PostgresConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"123456"`
	Database string `env:"DATABASE" envDefault:"depublic"`
}

// NewConfig creates a new config
func NewConfig(envPath string) (*Config, error) {
	cfg, err := parseConfig(envPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func parseConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, errors.New("failed to load env")
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return nil, errors.New("failed to parse config")
	}
	return cfg, nil
}
