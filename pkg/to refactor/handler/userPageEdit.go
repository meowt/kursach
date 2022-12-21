package handler

import (
	"Diploma/pkg/database"
	error2 "Diploma/pkg/error"
	"Diploma/pkg/to refactor/auth"
	"Diploma/server"
	"github.com/gorilla/mux"
	"net/http"
)

func UserPageEdit(w http.ResponseWriter, r *http.Request) {
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
	error2.errorProc(w, e, "Getting user page data error")

	//Add edit page...
	//...
}
