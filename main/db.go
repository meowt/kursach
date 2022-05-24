package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jmoiron/sqlx"
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

func dbRequestLogin(d struct{ email, password string }) (bool, user) {
	var e error
	var CurrentUser user
	sqlQuery := "SELECT password FROM users WHERE email = '" + d.email + "';"
	e = db.QueryRow(sqlQuery).Scan(&CurrentUser.Password)
	if e != nil {
		return false, CurrentUser
	}
	if comparePassword(d.password, CurrentUser.Password) {
		sqlQuery = "SELECT * FROM users WHERE email = '" + d.email + "';"
		e = db.QueryRow(sqlQuery).Scan(&CurrentUser.ID, &CurrentUser.Username, &CurrentUser.Email, &CurrentUser.Password)
		return true, CurrentUser
	} else {
		return false, CurrentUser
	}
}
