package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"Kursach/pkg/database"
)

var store = sessions.NewCookieStore([]byte("random-hash-secret"))

func Server() {
	rtr := mux.NewRouter()
	http.Handle("/", rtr)

	//Set assets handler function
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/"))))
	http.Handle("/user_files/", http.StripPrefix("/user_files/", http.FileServer(http.Dir("./web/user_files/"))))

	rtr.HandleFunc("/", Index)
	rtr.HandleFunc("/catalogue", Catalogue)

	//user page handler
	rtr.HandleFunc("/user/{username}", UserPage)
	rtr.HandleFunc("/user/{username}/follows", UserPageFollows)
	rtr.HandleFunc("/user/{username}/about", UserPageAbout)
	rtr.HandleFunc("/user/{username}/edit", UserPageEdit)

	//theme page handler
	rtr.HandleFunc("/theme/{id}", ThemePage)

	//upload & edit page handler
	rtr.HandleFunc("/upload", ThemeUploadPage)
	rtr.HandleFunc("/edit", profEdit)

	//post pages handler
	rtr.HandleFunc("/posts/login", Login).Methods("POST")
	rtr.HandleFunc("/posts/reg", Reg).Methods("POST")
	rtr.HandleFunc("/posts/upload", Upload).Methods("POST")
	rtr.HandleFunc("/posts/avatar", avatar).Methods("POST")
	rtr.HandleFunc("/posts/background", background).Methods("POST")
	rtr.HandleFunc("/posts/themeUpdate", themeUpdate).Methods("POST")

	//exit page handler
	rtr.HandleFunc("/exit", Exit).Methods("GET")
	//error page handler
	rtr.HandleFunc("/error", ErrorPage).Methods("GET")

	e := http.ListenAndServe(database.Cfg.ServerHost+":"+database.Cfg.ServerPort, context.ClearHandler(http.DefaultServeMux))
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось запустить сервер")
	}
}
