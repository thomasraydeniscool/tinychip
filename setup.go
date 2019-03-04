package main

import (
	"database/sql"

	"github.com/spf13/cobra"
)

func setup() *cobra.Command {
	return &cobra.Command{
		Use: "setup",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := setupDb(); err != nil {
				return err
			}
			return nil
		},
	}
}

func setupDb() error {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/")
	if err != nil {
		return err
	}
	db.Exec("CREATE DATABASE IF NOT EXISTS tinychip;")
	db.Close()
	return nil
}
