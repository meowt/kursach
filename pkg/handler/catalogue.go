package handler

import (
	"Kursach/pkg/database"
	"fmt"
	"html/template"
	"net/http"
)

func Catalogue(w http.ResponseWriter, r *http.Request) {
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

	lastThemes, e := database.GetCatalogue("LAST")
	errorProc(w, e, "Getting user page data error")

	meowtThemes, e := database.GetCatalogue("MEOWT")
	errorProc(w, e, "Getting user page data error")

	//Parsing templates
	t, e := template.ParseFiles(
		"./web/templates/scripts.html",
		"./web/templates/trueHeader.html",
		"./web/templates/catalogue.html",
		"./web/templates/catMeowt.html",
		"./web/templates/catHead.html")
	errorProc(w, e, "Template parsing error")

	//Executing templates with db data
	var headerData struct {
		Username string
	}
	headerData.Username = fmt.Sprint(session.Values["username"])

	e = t.ExecuteTemplate(w, "trueHeader", headerData)
	errorProc(w, e, "Template executing error")

	e = t.ExecuteTemplate(w, "catHead", nil)
	errorProc(w, e, "Template executing error")

	e = t.ExecuteTemplate(w, "catalogue", lastThemes)
	errorProc(w, e, "Template executing error")

	e = t.ExecuteTemplate(w, "catMeowt", meowtThemes)
	errorProc(w, e, "Template executing error")

	e = t.ExecuteTemplate(w, "scripts", nil)
	errorProc(w, e, "Template executing error")
}
