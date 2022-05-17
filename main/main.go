package main

import (
	"fmt"
	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

func main() {

	e := connect()
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	server()
}
