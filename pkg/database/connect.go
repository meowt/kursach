package database

import (
	"fmt"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var db *sqlx.DB

// Connect opens a database and pings it to prove connection.
// Function uses Cfg data that stores in pkg/database/dbSettings.cfg as a connection data.
func Connect() error {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Cfg.PgHost, Cfg.PgPort, Cfg.PgUser, Cfg.PgPass, Cfg.PgBase))
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	log.Println("Successfully connected to Postgres")
	return nil
}
