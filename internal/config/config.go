package config

import (
	"errors"
	"os"
	"github.com/joho/godotenv"
)

type MidtransConfig struct {
	ClientKey string
	ServerKey string
	// Other Midtrans configuration fields
}

type Config struct {
	Port           string         `env:"PORT" envDefault:"8080"`
	Postgres       PostgresConfig `envPrefix:"POSTGRES_"`
	JWT            JwtConfig      `envPrefix:"JWT_"`
	Session        SessionConfig  `envPrefix:"SESSION_"`
	Env            string
	MidtransConfig MidtransConfig
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
	// Set Midtrans configuration
	cfg.MidtransConfig.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	cfg.MidtransConfig.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")

	return cfg, nil
}

func parseConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, errors.New("failed to load env")
	}

	cfg := &Config{}
	err = parseEnv(cfg)
	if err != nil {
		return nil, errors.New("failed to parse config")
	}
	return cfg, nil
}

func parseEnv(cfg *Config) error {
	// You need to implement logic to parse environment variables
	// and set values to the appropriate fields in the Config struct.
	// Here is an example of how you might do it:

	cfg.Port = os.Getenv("PORT")
	cfg.JWT.SecretKey = os.Getenv("JWT_SECRET_KEY")
	cfg.Session.SecretKey = os.Getenv("SESSION_SECRET_KEY")

	// Similarly, parse and set PostgresConfig
	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")
	cfg.Postgres.User = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.Postgres.Database = os.Getenv("POSTGRES_DATABASE")

	return nil
}
