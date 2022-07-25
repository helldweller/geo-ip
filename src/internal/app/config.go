package app

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config is a structure containing configuration fields for this application.
type Config struct {
	Loglevel         string `env:"LOG_LEVEL"                   env-default:"error"`
	HTTPListenIPPort string `env:"HTTP_LISTEN"                 env-default:":80"`
}

var conf Config

func init() {
	err := cleanenv.ReadEnv(&conf)
	if err != nil {
		fmt.Printf("Something went wrong while reading the configuration: %s", err)
		os.Exit(1)
	}
}
