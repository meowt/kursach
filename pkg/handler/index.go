package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//Session start
	session, e := store.Get(r, "session-name")
	errorProc(w, e, "Session start error")

	//Parsing templates
	t, e := template.ParseFiles(
		"./web/templates/indexBody.html",
		"./web/templates/footer.html",
		"./web/templates/index.html",
		"./web/templates/header.html",
		"./web/templates/trueHeader.html",
		"./web/templates/scripts.html")
	errorProc(w, e, "Template parsing error")

	//Executing templates
	if AuthCheck(session) {
		var data struct {
			Id, Username string
		}
		data.Username = fmt.Sprint(session.Values["username"])

		e = t.ExecuteTemplate(w, "trueHeader", data)
		errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "loggedIndex", data)
		errorProc(w, e, "Template executing error")
	} else {
		e = t.ExecuteTemplate(w, "index", nil)
		errorProc(w, e, "Template executing error")
	}
}
