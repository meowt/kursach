package main

import (
	"fmt"
	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

func main() {
	e := connect()
	defer db.Close()
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	server()
}
