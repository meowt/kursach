package auth

import (
	"Diploma/pkg/database"
	error2 "Diploma/pkg/error"
	"Diploma/pkg/handler"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"io"
	"net/http"
	"os"
)

func Reg(w http.ResponseWriter, r *http.Request) {
	//Creating struct from POST data
	var RegData database.User
	RegData.Email = r.FormValue("email")
	RegData.Username = r.FormValue("username")
	RegData.Password = r.FormValue("password")

	//Uploading to DB
	var Message struct {
		Message string
	}
	Message.Message = RegData.RegRequest()
	fmt.Println(Message)

	if Message.Message != "" {
		//Error part
		//Parsing files
		t, err := template.ParseFiles("./web/templates/loginError.html")
		error2.errorProc(w, err, "Parsing files error")

		//Showing actual error
		err = t.ExecuteTemplate(w, "loginError", Message)
		error2.errorProc(w, err, "Executing templates error")
	} else {
		//Correct part
		src := "./web/user_files/avatar.jpg"
		dst := fmt.Sprintf("./web/user_files/%s/avatar.jpg", RegData.Username)
		in, err := os.Open(src)
		if err != nil {
			fmt.Println("error opening file")
		}
		defer in.Close()

		path := fmt.Sprintf("./web/user_files/%s", RegData.Username)
		err = os.MkdirAll(path, 0777)
		if err != nil {
			panic(err)
		}

		out, err := os.Create(dst)
		if err != nil {
			fmt.Println("error creating file")
		}
		defer out.Close()

		_, err = io.Copy(out, in)
		if err != nil {
			fmt.Println("error coping file")
		}
		src = "./web/user_files/background.jpg"
		dst = fmt.Sprintf("./web/user_files/%s/background.jpg", RegData.Username)
		in, err = os.Open(src)
		if err != nil {
			fmt.Println("error opening file")
		}
		defer in.Close()

		out, err = os.Create(dst)
		if err != nil {
			fmt.Println("error creating file")
		}
		defer out.Close()

		_, err = io.Copy(out, in)
		if err != nil {
			fmt.Println("error coping file")
		}

		//Session start
		session, e := handler.store.Get(r, "session-name")
		error2.errorProc(w, e, "Session start error")

		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   60 * 60 * 2,
			HttpOnly: true,
		}

		session.Values["logged"] = true
		session.Values["username"] = RegData.Username

		e = session.Save(r, w)
		error2.errorProc(w, e, "Session save error")

		http.Redirect(w, r, "/user/"+RegData.Username, http.StatusSeeOther)
	}
}
