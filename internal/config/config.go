package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	PostgresConfig struct {
		Host     string `yaml:"host" env:"POSTGRES_HOST" env-default:"localhost"`
		User     string `yaml:"user" env:"POSTGRES_USER" env-default:"postgres"`
		Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-default:"postgres"`
		Dbname   string `yaml:"dbname" env:"POSTGRES_DBNAME" env-default:"postgres"`
		Port     string `yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
		SSLMode  string `yaml:"sslmode" env:"POSTGRES_SSLMODE" env-default:"disable"`
		Timezone string `yaml:"timezone" env:"POSTGRES_TIMEZONE" env-default:"Europe/Moscow"`
	} `yaml:"postgres"`
}

func Load(path string) (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
