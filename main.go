package main

import (
	"fmt"
	"html"
	"net/http"

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

	http.HandleFunc("/graphql", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello, %q", html.EscapeString(req.URL.Path))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
}
