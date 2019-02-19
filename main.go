package main

import (
	"fmt"

	"github.com/caarlos0/env"
)

// Config for Environment variables
type Config struct {
	Home         string `env:"HOME"`
	Port         int    `env:"PORT" envDefault:"3000"`
	IsProduction bool   `env:"PRODUCTION"`
}

func main() {
	cfg := Config{}
	env.Parse(&cfg)
	fmt.Println(cfg)
}
