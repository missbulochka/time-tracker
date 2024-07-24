package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Env           string `envconfig:"TTRACKER_ENV" default:"dev"`
	MigrationPATH string `envconfig:"TTRACKER_MIGRATIONS_PATH" default:"./migrations"`
	HTTPcfg       HTTPConfig
	PSQLcfg       PSQLConfig
}

type HTTPConfig struct {
	HTTPServer string `envconfig:"TTRACKER_HTTP_SERVER" default:"0.0.0.0"`
	HTTPPort   string `envconfig:"TTRACKER_HTTP_PORT" default:"8080"`
}

type PSQLConfig struct {
	UserName string `envconfig:"TTRACKER_USER_NAME" default:"postgres"`
	Password string `envconfig:"TTRACKER_PASSWORD" default:"pass1234"`
	Host     string `envconfig:"TTRACKER_DB_HOST" default:"time-tracker-psql"`
	Port     string `envconfig:"TTRACKER_DB_PORT" default:"5432"`
	DBName   string `envconfig:"TTRACKER_DB_NAME" default:"Tracking"`
}

func LoadCfg() (*Config, error) {
	var cfg = new(Config)

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
