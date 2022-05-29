package main

import (
	"database/sql"
	"errors"
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
		e = db.QueryRow(sqlQuery).Scan(&CurrentUser.Username, &CurrentUser.Email, &CurrentUser.Password, &CurrentUser.Description, &CurrentUser.ID)
		return true, CurrentUser
	} else {
		return false, CurrentUser
	}
}

func dbRequestReg(d struct{ email, username, password string }) error {
	//checking of existing users with this email
	var userID string
	db.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE 'email' = %s  ;", d.email)).Scan(&userID)
	e := errors.New("Пользователь с этой почтой уже зарегистрирован")
	if userID != "" {
		return e
	}
	//inserting new user into db
	d.password, e = hashPassword(d.password)
	if e != nil {
		return e
	}
	sqlQuery := "INSERT INTO users (username, email, password) VALUES ('" + d.username + "', '" + d.email + "', '" + d.password + "');"
	_, e = db.Exec(sqlQuery)
	return e
}

func getUserPage(d string) error {
	var pageOwner user
	e := db.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE 'username' = %s  ;", d)).Scan(&pageOwner.Username, &pageOwner.Email, &pageOwner.Password, &pageOwner.Description, &pageOwner.ID)
	if e != nil {
		return e
	}
	return e
}
