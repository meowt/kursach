package auth

import (
	error2 "Diploma/pkg/error"
	"Diploma/pkg/handler"
	"github.com/gorilla/sessions"
	"net/http"
)

func Exit(w http.ResponseWriter, r *http.Request) {
	//Session start
	session, e := handler.store.Get(r, "session-name")
	error2.errorProc(w, e, "Session start error")

	//Session delete
	session.Values["username"] = ""
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	e = session.Save(r, w)
	error2.errorProc(w, e, "Session save error")

	//Redirecting to index page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
