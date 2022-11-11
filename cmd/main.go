package main

import (
	"Kursach/pkg/database"
	"Kursach/pkg/handler"
	"fmt"
)

func main() {
	//Connecting to Postgres
	e := database.Connect()
	if e != nil {
		fmt.Println(e.Error() + "\n Db connect error")
	}

	//Starting handling http requests
	handler.Server()

}
