package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type config struct {
	Port string `env:"PORT,required"`
	Path string `env:"PATH"`
}

var Cfg config

func LoadEnvConfig() error {
	godotenv.Load(".env")
	Cfg = config{}
	err := env.Parse(&Cfg)
	if err != nil {
		return err
	}

	return nil
}
