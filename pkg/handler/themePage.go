package handler

import (
	"Kursach/pkg/database"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func ThemePage(w http.ResponseWriter, r *http.Request) {
	//Session start
	session, e := store.Get(r, "session-name")
	errorProc(w, e, "Session start error")

	//Session expiring update
	if AuthCheck(session) {
		UpdateSession(session, w, r)
	} else {
		//Redirecting not auth users
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	//Parsing url
	vars := mux.Vars(r)

	//Getting theme's data
	var theme database.Theme
	err := theme.GetByID(vars["id"])
	errorProc(w, err, "Getting theme's data error")
	//Parsing templates
	t, e := template.ParseFiles(
		"./web/templates/scripts.html",
		"./web/templates/trueHeader.html",
		"./web/templates/themePage.html")
	errorProc(w, e, "Template parsing error")

	//Executing templates with db data
	var headerData struct {
		Username string
	}
	headerData.Username = fmt.Sprint(session.Values["username"])
	if vars["username"] == fmt.Sprint(session.Values["username"]) {
		//Own page
		e = t.ExecuteTemplate(w, "trueHeader", headerData)
		errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "themeOwnPage", theme)
		errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "scripts", nil)
		errorProc(w, e, "Template executing error")
	} else {
		//Else's page
		e = t.ExecuteTemplate(w, "trueHeader", headerData)
		errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "themeOwnPage", theme)
		errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "scripts", nil)
		errorProc(w, e, "Template executing error")
	}
}
