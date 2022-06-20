package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	//Session start
	session, e := store.Get(r, "session-name")
	errorProc(w, e, "Session start error")

	//Parsing templates
	t, e := template.ParseFiles(
		"./web/templates/header.html",
		"./web/templates/trueHeader.html",
		"./web/templates/scripts.html",
		"./web/templates/error.html")
	errorProc(w, e, "Template parsing error")

	//Executing templates
	var headerData struct {
		Username string
	}
	if session.Values["username"] != "" {
		headerData.Username = fmt.Sprint(session.Values["username"])
		e = t.ExecuteTemplate(w, "trueHeader", headerData)
		errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "error", nil)
		errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "scripts", nil)
		errorProc(w, e, "Template executing error")
	} else {
		e = t.ExecuteTemplate(w, "header", nil)
		errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "error", nil)
		errorProc(w, e, "Template executing error")
	}
}
