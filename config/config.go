package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Env 		string			`env:"ENV" evDefault:"dev"`
	Host    	string	        `env:"HOST" envDefault:"localhost"`
	Port		string			`env:"PORT" envDefault: "8080"`
	Postgres    PostgresConfig	`envPrefix:"POSTGRES_"`
}

type PostgresConfig struct {
	Host 		string 			`env:"HOST" envDefault:"localhost"`
	Port		string			`env:"PORT" envDefault: "5432"`
	User		string			`env:"USER" envDefault: "postgres"`
	Password	string			`env:"PASSWORD" envDefault: "postgres"`
	Database	string			`env:"DATABASE" envDefault: "postgres"`
}

func NewConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	err = env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}