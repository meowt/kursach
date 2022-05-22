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

func dbRequestLogin(d string) (bool, user) {
	var e error
	var OneUser user
	e = db.QueryRow(
		"SELECT * FROM users WHERE email = '"+d+"';",
	).Scan(&OneUser.ID, &OneUser.Username, &OneUser.Email)
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось отправить запрос к БД для авторизации")
	} else {
		return true, OneUser
	}
}
