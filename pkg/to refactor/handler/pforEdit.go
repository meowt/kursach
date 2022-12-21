package handler

import (
	error2 "Diploma/pkg/error"
	"Diploma/pkg/to refactor/auth"
	"Diploma/server"
	"fmt"
	"html/template"
	"net/http"
)

func profEdit(w http.ResponseWriter, r *http.Request) {
	//Session start
	session, e := server.store.Get(r, "session-name")
	error2.errorProc(w, e, "Session start error")

	//Session expiring update
	if auth.AuthCheck(session) {
		auth.UpdateSession(session, w, r)
	} else {
		//Redirecting not auth users
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	//Parsing templates
	t, e := template.ParseFiles(
		"./web/templates/scripts.html",
		"./web/templates/trueHeader.html",
		"./web/templates/profEditPage.html")
	error2.errorProc(w, e, "Template parsing error")

	//Executing templates with db data
	var headerData struct {
		Username string
	}
	headerData.Username = fmt.Sprint(session.Values["username"])

	e = t.ExecuteTemplate(w, "trueHeader", headerData)
	error2.errorProc(w, e, "Template executing error")

	e = t.ExecuteTemplate(w, "profEditPage", nil)
	error2.errorProc(w, e, "Template executing error")

	e = t.ExecuteTemplate(w, "scripts", nil)
	error2.errorProc(w, e, "Template executing error")
}
