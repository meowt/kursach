package database

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() error {
	var e error
	db, e = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Cfg.PgHost, Cfg.PgPort, Cfg.PgUser, Cfg.PgPass, Cfg.PgBase))
	if e != nil {
		return e
	}
	return nil
}
