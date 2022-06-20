package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func ThemeUploadPage(w http.ResponseWriter, r *http.Request) {
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

	//Parsing templates
	t, e := template.ParseFiles(
		"./web/templates/scripts.html",
		"./web/templates/trueHeader.html",
		"./web/templates/uploadPage.html")
	errorProc(w, e, "Template parsing error")

	//Executing templates with db data
	var headerData struct {
		Username string
	}
	headerData.Username = fmt.Sprint(session.Values["username"])

	e = t.ExecuteTemplate(w, "trueHeader", headerData)
	errorProc(w, e, "Template executing error")

	e = t.ExecuteTemplate(w, "uploadPage", nil)
	errorProc(w, e, "Template executing error")

	e = t.ExecuteTemplate(w, "scripts", nil)
	errorProc(w, e, "Template executing error")
}
