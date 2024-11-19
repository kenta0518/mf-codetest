package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		MySQL    `yaml:"mysql"`
		Settings `yaml:"settings"`
	}

	MySQL struct {
		DBConn string `env-required:"true" yaml:"db_conn" env:"DB_CONN"`
	}

	Settings struct {
		Environment    string `yaml:"environment" env:"SETTING_ENVIRONMENT"`
		ResourceServer string `yaml:"resource_server" env:"SETTING_RESOURCE_SERVER"`
	}
)

func (s Settings) IsDevelopment() bool {
	return s.Environment == "Development"
}

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
