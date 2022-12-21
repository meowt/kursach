package database

import (
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func Setup() (PostgresClient *sqlx.DB, err error) {
	PostgresClient, err = Connect()
	if err != nil {
		return
	}
	if err = Deploy(PostgresClient); err != nil {
		return
	}
	return
}

// Connect opens database and pings it to prove connection.
// Function uses Cfg data that stores in pkg/config/config.yml as a connection data.
func Connect() (PostgresClient *sqlx.DB, err error) {
	PostgresClient, err = sqlx.Open("postgres", viper.GetString("postgres.postgresDsn"))
	if err != nil {
		return
	}
	if err = PostgresClient.Ping(); err != nil {
		return
	}
	log.Println("Successfully connected to Postgres")
	return
}

func Deploy(PostgresClient *sqlx.DB) (err error) {
	for _, command := range viper.GetStringMapString("postgres.deployment") {
		if _, err = PostgresClient.Exec(command); err != nil {
			return
		}
		log.Printf("Success: %s\n", command)
	}
	log.Println("Successfully deployed database")
	return
}
