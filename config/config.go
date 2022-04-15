package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	DatabaseURI            string `env:"DATABASE_URI" envDefault:"mongodb://localhost:27017/"`
	Database               string `env:"DATABASE" envDefault:"DEFAULT_DATABASE.db"`
	DatabaseUserCollection string `env:"USER_COLLECTION" envDefault:"USER_COLLECTION"`
	Environment            string `env:"ENVIRONMENT" envDefault:"DEVELOPMENT"`
	Port                   string `env:"PORT" envDefault:"8001"`
	AppName                string `env:"APP_NAME" envDefault:"TODO_API"`
	JWTExpirationMs        string `env:"JWT_EXPIRATION_MS" envDefault:"86400000"`
	JWTSecret              string `env:"JWT_SECRET" envDefault:"dwQul1cohWnvkofm7kDtemQ7gSGT-c5ns4uLaOWJ2Gs"`
}

func New() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(fmt.Sprintf(`Error on parse env variable due error:[%v]`, err))
	}

	return &cfg
}
