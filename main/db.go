package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func connect() error {
	var e error
	db, e = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgBase))
	if e != nil {
		return e
	}

	return nil
}
