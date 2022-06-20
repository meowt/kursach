package pkg

//
//import (
//	"database/sql"
//	"errors"
//	"fmt"
//	_ "github.com/jmoiron/sqlx"
//)
////var db *sql.DB
//
//func connect() error {
//	var e error
//	db, e = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Cfg.PgHost, Cfg.PgPort, Cfg.PgUser, Cfg.PgPass, Cfg.PgBase))
//	if e != nil {
//		return e
//	}
//	defer db.Close()
//	return nil
//}
//
//func dbRequestLogin(d struct{ email, password string }) (bool, User) {
//	var CurrentUser User
//	sqlQuery := "SELECT password FROM users WHERE email = '" + d.email + "';"
//	e := db.QueryRow(sqlQuery).Scan(&CurrentUser.Password)
//	if e != nil {
//		return false, CurrentUser
//	}
//	if comparePassword(d.password, CurrentUser.Password) {
//		sqlQuery = "SELECT * FROM users WHERE email = '" + d.email + "';"
//		e = db.QueryRow(sqlQuery).Scan(&CurrentUser.Username, &CurrentUser.Email, &CurrentUser.Password, &CurrentUser.Description, &CurrentUser.ID)
//		return true, CurrentUser
//	} else {
//		return false, CurrentUser
//	}
//}
//
//func dbRequestReg(d struct{ email, username, password string }) error {
//	//checking of existing users with this email
//	var userID string
//	e := db.QueryRow(fmt.Sprintf("SELECT * FROM users WHERE 'email' = %s  ;", d.email)).Scan(&userID)
//	e = errors.New("пользователь с этой почтой уже зарегистрирован")
//	if userID != "" {
//		return e
//	}
//	//inserting new user into db
//	d.password, e = hashPassword(d.password)
//	if e != nil {
//		return e
//	}
//	sqlQuery := "INSERT INTO users (username, email, password) VALUES ('" + d.username + "', '" + d.email + "', '" + d.password + "');"
//	_, e = db.Exec(sqlQuery)
//	return e
//}
//
//func getThemePage(d string) Theme {
//	//getting theme data
//	var theme Theme
//	e := db.QueryRow(fmt.Sprintf("SELECT * FROM themes WHERE id = '%s' ;", d)).Scan(
//		&theme.Path,
//		&theme.ReleaseDate,
//		&theme.CreatorName,
//		&theme.Followers,
//		&theme.Description,
//		&theme.ID,
//		&theme.Name)
//	if e != nil {
//		fmt.Println("theme scanning" + e.Error())
//	}
//	return theme
//}
//
//func getThemeId(d string) int {
//	var id int
//	_ = db.QueryRow(fmt.Sprintf("SELECT id FROM themes WHERE creator_name = '%s' ORDER BY id DESC LIMIT 1;", d)).Scan(&id)
//	return id
//}
//
//func saveTheme(t Theme) error {
//	//inserting new user into db
//	sqlQuery := "INSERT INTO themes (name, path, creator_name, description) VALUES ('" + t.Name + "', '" + t.Path + "', '" + t.CreatorName + "', '" + t.Description.String + "');"
//	_, e := db.Exec(sqlQuery)
//	return e
//}
