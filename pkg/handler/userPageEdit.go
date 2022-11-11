package handler

import (
	"Kursach/pkg/database"
	"github.com/gorilla/mux"
	"net/http"
)

func UserPageEdit(w http.ResponseWriter, r *http.Request) {
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

	//Add edit page...
	//...
}
