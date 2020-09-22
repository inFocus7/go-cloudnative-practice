// Imports and sets up config file structure for usage elsewhere.
package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

type Conf struct {
	Server serverConf
	Debug  bool `required:"true"`
	Db     dbConf
}

type serverConf struct {
	Port         int           `split_words:"true" required:"true"`
	TimeoutRead  time.Duration `split_words:"true" required:"true"`
	TimeoutWrite time.Duration `split_words:"true" required:"true"`
	TimeoutIdle  time.Duration `split_words:"true" required:"true"`
}

type dbConf struct {
	Host     string `split_words:"true" required:"true"`
	Port     int    `split_words:"true" required:"true"`
	Username string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASS" required:"true"`
	Name     string `split_words:"true" required:"true"`
}

func AppConfig() *Conf {
	var c Conf

	if err := envconfig.Process("", &c); err != nil {
		log.Fatalf("Failed to process env: %v", err)
	}

	return &c
}
