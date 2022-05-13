package main

//
//import (
//	"bytes"
//	"fmt"
//	"html/template"
//)
//
//func starter() {
//	buffer := new(bytes.Buffer)
//
//	templateParser, e := template.ParseFiles("./web/templates/index.html")
//	if e != nil {
//		fmt.Println(e.Error())
//		panic("Не удалось открыть index.html")
//	}
//
//	e = templateParser.ExecuteTemplate(buffer, "index", map[string]string{})
//
//	if e != nil {
//		fmt.Println(e.Error())
//		panic("Не удалось запустить index.html")
//	}
//	//fmt.Println(buffer.String())
//
//}
