package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jmoiron/sqlx"
)

var db *sql.DB

type UsersIdUsername struct {
	id       []string
	username []string
}

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

func dbRequestReg(d struct{ email, password string }) (sql.Result, error) {
	//checking of existing users with this email
	sqlQuery := "SELECT * FROM users WHERE 'email' = ?;"
	row := db.QueryRow(sqlQuery, d.email)
	var userID string
	e := row.Scan(&userID)
	if e != sql.ErrNoRows {
		return nil, e
	}
	//inserting new user into db
	d.password, e = hashPassword(d.password)
	if e != nil {
		return nil, e
	}
	sqlQuery = "INSERT INTO users (email, password) VALUES ('" + d.email + "', '" + d.password + "');"
	res, e := db.Exec(sqlQuery)
	return res, e
}

func receiveAllUsersID() (struct {
	id       []string
	username []string
}, error) {
	var users UsersIdUsername
	sqlQuery := "SELECT id, username FROM users;"
	rows, e := db.Query(sqlQuery)
	if e != nil {
		return users, e
	}
	rows.Scan(&users.id, &users.username)
	return users, e
}
