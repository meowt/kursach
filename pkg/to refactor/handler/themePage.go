package handler

import (
	"Diploma/pkg/database"
	error2 "Diploma/pkg/error"
	"Diploma/pkg/to refactor/auth"
	"Diploma/server"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func ThemePage(w http.ResponseWriter, r *http.Request) {
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

	//Parsing url
	vars := mux.Vars(r)

	//Getting theme's data
	var theme database.Theme
	err := theme.GetByID(vars["id"])
	error2.errorProc(w, err, "Getting theme's data error")
	//Parsing templates
	t, e := template.ParseFiles(
		"./web/templates/scripts.html",
		"./web/templates/trueHeader.html",
		"./web/templates/themePage.html")
	error2.errorProc(w, e, "Template parsing error")

	//Executing templates with db data
	var headerData struct {
		Username string
	}
	headerData.Username = fmt.Sprint(session.Values["username"])
	if vars["username"] == fmt.Sprint(session.Values["username"]) {
		//Own page
		e = t.ExecuteTemplate(w, "trueHeader", headerData)
		error2.errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "themeOwnPage", theme)
		error2.errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "scripts", nil)
		error2.errorProc(w, e, "Template executing error")
	} else {
		//Else's page
		e = t.ExecuteTemplate(w, "trueHeader", headerData)
		error2.errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "themeOwnPage", theme)
		error2.errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "scripts", nil)
		error2.errorProc(w, e, "Template executing error")
	}
}
