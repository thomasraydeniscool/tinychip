package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

type config struct {
	Port       int    `env:"PORT" envDefault:"3000"`
	Db         string `env:"DB"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
}

// Order struct
type Order struct {
	Total float64 `json:"total"`
}

// GetOrders returns orders
func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []Order
	json.NewEncoder(w).Encode(orders)
}

func start() *cobra.Command {
	return &cobra.Command{
		Use: "start",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config{}

			if err := env.Parse(&cfg); err != nil {
				return err
			}

			addr := fmt.Sprintf(":%d", cfg.Port)

			r := mux.NewRouter()

			r.HandleFunc("/orders", GetOrders).Methods("GET")

			log.Println("Listening on", addr)

			if err := http.ListenAndServe(addr, r); err != nil {
				return err
			}

			return nil
		},
	}
}

func main() {
	cmd := &cobra.Command{
		Use:   "tinychip",
		Short: "Tiny headless eCommerce platform",
	}

	cmd.AddCommand(setup())
	cmd.AddCommand(start())

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
