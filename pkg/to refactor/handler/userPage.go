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

func UserPage(w http.ResponseWriter, r *http.Request) {
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

	//Getting info about current page
	var pageOwner database.User
	e = pageOwner.GetPageData(vars["username"])
	if e != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}

	themes, e := database.GetCreatorsThemes(vars["username"])
	error2.errorProc(w, e, "Getting user page data error")

	//Parsing templates
	t, e := template.ParseFiles(
		"./web/templates/scripts.html",
		"./web/templates/trueHeader.html",
		"./web/templates/userPage.html",
		"./web/templates/userPageHeadThemes.html")
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

		e = t.ExecuteTemplate(w, "userOwnPageHead", pageOwner)
		error2.errorProc(w, e, "Template executing error")

		if len(themes) != 0 {
			e = t.ExecuteTemplate(w, "userPageTheme", themes)
			error2.errorProc(w, e, "Template executing error")
		} else {
			e = t.ExecuteTemplate(w, "emptyUserPageTheme", themes)
			error2.errorProc(w, e, "Template executing error")
		}

		e = t.ExecuteTemplate(w, "scripts", nil)
		error2.errorProc(w, e, "Template executing error")
	} else {
		//Else's page
		e = t.ExecuteTemplate(w, "trueHeader", headerData)
		error2.errorProc(w, e, "Template executing error")

		e = t.ExecuteTemplate(w, "userPageHead", pageOwner)
		error2.errorProc(w, e, "Template executing error")

		if len(themes) != 0 {
			e = t.ExecuteTemplate(w, "userPageTheme", themes)
			error2.errorProc(w, e, "Template executing error")
		} else {
			e = t.ExecuteTemplate(w, "emptyUserPageTheme", themes)
			error2.errorProc(w, e, "Template executing error")
		}

		e = t.ExecuteTemplate(w, "scripts", nil)
		error2.errorProc(w, e, "Template executing error")
	}
}
