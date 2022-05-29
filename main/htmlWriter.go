package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

var router *gin.Engine
var data user
var store = sessions.NewCookieStore([]byte("random-hash-secret"))

func server() {
	rtr := mux.NewRouter()
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/"))))
	rtr.HandleFunc("/", index)
	http.Handle("/", rtr)
	rtr.HandleFunc("/user/{username}", userPage)
	rtr.HandleFunc("/posts/login", login).Methods("POST")
	rtr.HandleFunc("/posts/reg", reg).Methods("POST")
	e := http.ListenAndServe(cfg.ServerHost+":"+cfg.ServerPort, context.ClearHandler(http.DefaultServeMux))
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось запустить сервер")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("./web/templates/indexBody.html", "./web/templates/header.html", "./web/templates/footer.html", "./web/templates/index.html", "./web/templates/trueHeader.html")

	if e != nil {
		fmt.Fprintln(w, e.Error())
	}
	var data struct {
		user string
	}
	data.user = "meowt"
	t.ExecuteTemplate(w, "index", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	//creating struct from POST data
	var LoginData struct {
		email, password string
	}
	LoginData.email = r.FormValue("email")
	LoginData.password = r.FormValue("password")
	//checking
	var res bool
	res, data = dbRequestLogin(LoginData)
	if res {

		//r.GET("/incr", func(c *gin.Context) {
		//	session := sessions.Default(c)
		//	var count int
		//	v := session.Get("count")
		//	if v == nil {
		//		count = 0
		//	} else {
		//		count = v.(int)
		//		count++
		//	}
		//	session.Set("count", count)
		//	session.Save()
		//	c.JSON(200, gin.H{"count": count})
		//})

		http.Redirect(w, r, "/user/"+data.Username, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func reg(w http.ResponseWriter, r *http.Request) {
	//creating struct from POST data
	var RegData struct {
		email, username, password string
	}
	RegData.email = r.FormValue("email")
	RegData.username = r.FormValue("username")
	RegData.password = r.FormValue("password")
	//uploading to db
	e := dbRequestReg(RegData)
	if e != nil {
		//error part
		http.Redirect(w, r, "/", http.StatusSeeOther)
		//c.Writer.WriteString("<script>" +
		//	"alert('Email is busy.')" +
		//	"</script>")
		//c.Writer.WriteString("<script>" +
		//	"window.location.href = 'http://127.0.0.1:9090/'" +
		//	"</script>")
	} else {
		//correct part
		http.Redirect(w, r, "/"+RegData.username, http.StatusSeeOther)
	}
}

func userPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	e := getUserPage(vars["username"])
	if e != nil {

	}
	t, e := template.ParseFiles("./web/templates/indexBody.html", "./web/templates/header.html", "./web/templates/footer.html", "./web/templates/index.html", "./web/templates/trueHeader.html")
	t.ExecuteTemplate(w, "userpage", data)
}
