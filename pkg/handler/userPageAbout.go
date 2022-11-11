package handler

import (
	"Kursach/pkg/database"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func UserPageAbout(w http.ResponseWriter, r *http.Request) {
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

	//Getting info about current page
	var pageOwner database.User
	e = pageOwner.GetPageData(vars["username"])
	errorProc(w, e, "Getting user page data error")

	//Parsing templates
	t, e := template.ParseFiles("./web/templates/scripts.html",
		"./web/templates/trueHeader.html",
		"./web/templates/userPage.html",
		"./web/templates/userPageHeadAbout.html")
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

		e = t.ExecuteTemplate(w, "userOwnPageHead", pageOwner)
		errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "userPageAbout", pageOwner)
		errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "scripts", nil)
		errorProc(w, e, "Template executing error")
	} else {
		//Else's page
		e = t.ExecuteTemplate(w, "trueHeader", headerData)
		errorProc(w, e, "Template executing error")

		if pageOwner.Description.Valid {
			//Not nil description
			e = t.ExecuteTemplate(w, "userPageHead", pageOwner)
			errorProc(w, e, "Template executing error")

			e = t.ExecuteTemplate(w, "userPageAbout", pageOwner)
			errorProc(w, e, "Template executing error")
		} else {
			//Nil description
			e = t.ExecuteTemplate(w, "userPageHead", pageOwner)
			errorProc(w, e, "Template executing error")

			e = t.ExecuteTemplate(w, "emptyUserPageAbout", pageOwner)
			errorProc(w, e, "Template executing error")
		}

		e = t.ExecuteTemplate(w, "scripts", nil)
		errorProc(w, e, "Template executing error")
	}
}
