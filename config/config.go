package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		MySQL `yaml:"mysql"`
	}

	MySQL struct {
		DBConn string `env-required:"true" yaml:"db_conn" env:"DB_CONN"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
