// Imports and sets up config file structure for usage elsewhere.
package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

type Conf struct {
	Server serverConf
}

type serverConf struct {
	Port         int           `split_words:"true" required:"true"`
	TimeoutRead  time.Duration `split_words:"true" required:"true"`
	TimeoutWrite time.Duration `split_words:"true" required:"true"`
	TimeoutIdle  time.Duration `split_words:"true" required:"true"`
}

func AppConfig() *Conf {
	var c Conf

	if err := envconfig.Process("", &c); err != nil {
		log.Fatalf("Failed to process env: %v", err)
	}

	return &c
}
