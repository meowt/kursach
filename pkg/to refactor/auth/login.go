package auth

import (
	"Diploma/pkg/database"
	error2 "Diploma/pkg/error"
	"Diploma/pkg/handler"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//Checking
	var userData database.User
	var Message struct {
		Message string
	}
	Message.Message = userData.LoginRequest(r.FormValue("email"), r.FormValue("password"))
	if Message.Message == "" {
		//Correct part
		session, e := handler.store.Get(r, "session-name")
		error2.errorProc(w, e, "Session start error")

		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   60 * 60 * 2,
			HttpOnly: true,
		}

		data := userData

		session.Values["logged"] = true
		session.Values["username"] = data.Username

		e = session.Save(r, w)
		error2.errorProc(w, e, "Session save error")
		fmt.Println("/user/" + data.Username)
		http.Redirect(w, r, "/user/"+data.Username, http.StatusOK)
	} else {
		fmt.Println(Message)
		//Error part
		t, e := template.ParseFiles("./web/templates/loginError.html")
		error2.errorProc(w, e, "Parsing files error")

		e = t.ExecuteTemplate(w, "loginError", Message)
		error2.errorProc(w, e, "Executing templates error")

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
