package main

import (
	"Kursach/pkg/database"
	"Kursach/pkg/handler"
	"fmt"
)

func main() {
	e := database.Connect()
	if e != nil {
		fmt.Println(e.Error() + "\n Db connect error")
	}
	handler.Server()
	fmt.Println("Main file")

}
