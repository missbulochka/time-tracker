package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Env     string `envconfig:"TTRACKER_ENV" default:"dev"`
	HTTPcfg HTTPConfig
}

type HTTPConfig struct {
	HTTPServer string `envconfig:"TTRACKER_HTTP_SERVER" default:"0.0.0.0"`
	HTTPPort   string `envconfig:"TTRACKER_HTTP_PORT" default:"8080"`
}

func LoadCfg() (*Config, error) {
	var cfg = new(Config)

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
