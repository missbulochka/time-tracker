package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Env     string `env:"TTRACKER_ENV, default=local"`
	HTTPcfg HTTPConfig
}

type HTTPConfig struct {
	HTTPServer string `env:"TTRACKER_HTTP_SERVER, default=0.0.0.0"`
	HTTPPort   string `env:"TTRACKER_HTTP_PORT, default=8080"`
}

func LoadCfg() (*Config, error) {
	var cfg = new(Config)

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
