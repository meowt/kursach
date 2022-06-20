package database

import (
	"Kursach/pkg/service"
	"database/sql"
)

func LoginRequest(email, password string) (User, string) {
	var CurrentUser User
	sqlQuery := "SELECT password FROM users WHERE email = '" + email + "';"
	e := db.QueryRow(sqlQuery).Scan(&CurrentUser.Password)
	if e != nil && e != sql.ErrNoRows {
		return CurrentUser, "Произошла ошибка при авторизации"
	}
	if e == sql.ErrNoRows {
		return CurrentUser, "Пользователь с такой почтой не зарегистрирован"
	}
	if service.ComparePassword(password, CurrentUser.Password) {
		sqlQuery = "SELECT * FROM users WHERE email = '" + email + "';"
		e = db.QueryRow(sqlQuery).Scan(&CurrentUser.Username, &CurrentUser.Email, &CurrentUser.Password, &CurrentUser.Description, &CurrentUser.ID)
		return CurrentUser, ""
	} else {
		return CurrentUser, "Неправильный пароль"
	}
}
