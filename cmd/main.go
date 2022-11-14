package main

import (
	"Kursach/pkg/database"
	"Kursach/pkg/handler"
	"Kursach/pkg/service"
	"log"
)

func main() {
	//Logging initialisation
	err := service.LogInit()
	if err != nil {
		log.Fatal(err.Error() + "\nLogging init error")
	}

	//Connecting to Postgres
	err = database.Connect()
	if err != nil {
		log.Fatal(err.Error() + "\nDb connect error")
	}

	//Starting handling http requests
	err = handler.Server()
	if err != nil {
		log.Fatal(err.Error() + "\nServer starting error")
	}
}
