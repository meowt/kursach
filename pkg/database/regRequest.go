package database

import (
	"Kursach/pkg/service"
	"database/sql"
	"fmt"
	"strconv"
)

type RegType struct {
	Email, Username, Password string
}

func RegRequest(d RegType) string {
	//Checking of existing users with this email

	var scanStruct User
	row := db.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE email = '%s' LIMIT 1;", d.Email))
	e := row.Scan(&scanStruct.Username, &scanStruct.Email, &scanStruct.Password, &scanStruct.Description, &scanStruct.ID)

	if e != nil && e != sql.ErrNoRows {
		return "error"
	}

	userID := strconv.Itoa(scanStruct.ID)

	if userID != "0" {
		return "Пользователь с такой почтой уже зарегистрирован"
	}

	//Checking of existing users with this email
	row = db.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE username = '%s' LIMIT 1;", d.Username))
	e = row.Scan(&scanStruct.Username, &scanStruct.Email, &scanStruct.Password, &scanStruct.Description, &scanStruct.ID)
	if e != nil && e != sql.ErrNoRows {
		return "error"
	}

	userName := scanStruct.Username

	if userName != "" {
		return "Пользователь с таким именем уже существует"
	}

	//Hashing password
	d.Password, e = service.HashPassword(d.Password)
	if e != nil {
		return "error"
	}

	//Inserting new user into db
	sqlQuery := "INSERT INTO users (username, email, password) VALUES ('" + d.Username + "', '" + d.Email + "', '" + d.Password + "');"
	_, e = db.Exec(sqlQuery)
	if e != nil {
		return "error"
	}
	return ""
}
