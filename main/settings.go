package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type settings struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPass     string
	PgBase     string
	Html       string
	Assets     string
}

type user struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

var cfg settings

func init() {
	file, e := os.Open("main/settings.cfg")
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось открыть файл конфигурации")
	}
	defer file.Close()

	stat, e := file.Stat()
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось прочитать информацию о файле конфигурации")
	}

	readByte := make([]byte, stat.Size())

	_, e = file.Read(readByte)
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось прочитать файл конфигурации")
	}

	e = json.Unmarshal(readByte, &cfg)
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось преобразовать файл конфигурации")
	}
}
