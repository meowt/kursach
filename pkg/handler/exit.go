package handler

import (
	"github.com/gorilla/sessions"
	"net/http"
)

func Exit(w http.ResponseWriter, r *http.Request) {
	//Session start
	session, e := store.Get(r, "session-name")
	errorProc(w, e, "Session start error")

	//Session delete
	session.Values["username"] = ""
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	e = session.Save(r, w)
	errorProc(w, e, "Session save error")

	//Redirecting to index page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
