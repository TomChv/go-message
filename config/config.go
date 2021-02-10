package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type config struct {
	Server struct {
		Port string `envconfig:"API_PORT"`
	}
	Database struct {
		Name     string `envconfig:"DB_NAME" required:"true"`
		Username string `envconfig:"DB_USER" required:"true"`
		Password string `envconfig:"DB_PASS" required:"true"`
		Host     string `envconfig:"DB_HOST" required:"true"`
		Port     string `envconfig:"DB_PORT" required:"true"`
	}
}

var Config config

func init() {
	if err := envconfig.Process("", &Config); err != nil {
		log.Fatal(err)
	}
}
