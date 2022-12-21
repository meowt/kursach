package auth

import (
	error2 "Diploma/pkg/error"
	"github.com/gorilla/sessions"
	"net/http"
)

func AuthCheck(s *sessions.Session) bool {
	if s.Values["logged"] == true {
		return true
	} else {
		return false
	}
}

func UpdateSession(s *sessions.Session, w http.ResponseWriter, r *http.Request) {
	s.Options = &sessions.Options{
		MaxAge: 60 * 60 * 2,
	}
	e := s.Save(r, w)
	error2.errorProc(w, e, "Session saving error")
}
