package database

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

var Cfg settings

func init() {
	file, e := os.Open("pkg/database/dbSettings.cfg")
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

	e = json.Unmarshal(readByte, &Cfg)
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось преобразовать файл конфигурации")
	}
}
