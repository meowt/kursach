package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var data user
var store = sessions.NewCookieStore([]byte("random-hash-secret"))

func server() {
	rtr := mux.NewRouter()
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/"))))
	http.Handle("/user_files/", http.StripPrefix("/user_files/", http.FileServer(http.Dir("./web/user_files/"))))
	rtr.HandleFunc("/", index)
	http.Handle("/", rtr)
	rtr.HandleFunc("/user/{username}", userPage)
	rtr.HandleFunc("/user/{username}/follows", userPageFollows)
	rtr.HandleFunc("/user/{username}/about", userPageAbout)
	rtr.HandleFunc("/user/{username}/edit", userPageEdit)
	rtr.HandleFunc("/theme/{id}", themePage)
	rtr.HandleFunc("/upload", themeUpload)
	rtr.HandleFunc("/posts/login", login).Methods("POST")
	rtr.HandleFunc("/posts/reg", reg).Methods("POST")
	rtr.HandleFunc("/posts/upload", upload).Methods("POST")
	rtr.HandleFunc("/exit", exit).Methods("GET")
	e := http.ListenAndServe(cfg.ServerHost+":"+cfg.ServerPort, context.ClearHandler(http.DefaultServeMux))
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось запустить сервер")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	t, e := template.ParseFiles("./web/templates/indexBody.html", "./web/templates/header.html", "./web/templates/footer.html", "./web/templates/index.html", "./web/templates/trueHeader.html", "./web/templates/scripts.html")

	if e != nil {
		fmt.Fprintln(w, e.Error())
	}
	var data struct {
		Id, Username string
	}

	session, _ := store.Get(r, "session-name")
	fmt.Println("Index data:", session.Values["logged"], session.Values["username"])
	if session.Values["logged"] == true {
		data.Username = fmt.Sprint(session.Values["username"])
		t.ExecuteTemplate(w, "trueHeader", data)
		t.ExecuteTemplate(w, "loggedIndex", data)
	} else {
		t.ExecuteTemplate(w, "index", nil)
	}
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
		session, e := store.Get(r, "session-name")
		if e != nil {
			fmt.Fprintln(w, "Session error")
		}
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   60 * 60 * 2,
			HttpOnly: true,
		}
		session.Values["logged"] = true
		session.Values["username"] = data.Username
		e = session.Save(r, w)
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/user/"+data.Username, http.StatusSeeOther)
	} else {
		t, e := template.ParseFiles("./web/templates/loginError.html")
		if e != nil {
			fmt.Fprintln(w, e.Error())
		}
		t.ExecuteTemplate(w, "loginError", nil)
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
	} else {
		//correct part
		http.Redirect(w, r, "/"+RegData.username, http.StatusSeeOther)
	}
}

func userPage(w http.ResponseWriter, r *http.Request) {
	//session expiring update
	session, _ := store.Get(r, "session-name")
	fmt.Println("User page:", session.Values["logged"], session.Values["username"])
	if session.Values["logged"] != true {
		//http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	session.Options = &sessions.Options{
		MaxAge: 60 * 60 * 2,
	}
	_ = session.Save(r, w)
	vars := mux.Vars(r)
	pageOwner, themes, _ := getUserPage(vars["username"])

	t, _ := template.ParseFiles(
		"./web/templates/scripts.html",
		"./web/templates/trueHeader.html",
		"./web/templates/userPage.html",
		"./web/templates/userPageHeadThemes.html")

	data.Username = fmt.Sprint(session.Values["username"])
	if vars["username"] == fmt.Sprint(session.Values["username"]) {
		//fmt.Println("1 theme")
		_ = t.ExecuteTemplate(w, "trueHeader", data)
		_ = t.ExecuteTemplate(w, "userOwnPageHead", data)
		_ = t.ExecuteTemplate(w, "userPageTheme", themes)
		_ = t.ExecuteTemplate(w, "scripts", nil)
	} else {
		//fmt.Println("2 theme")
		_ = t.ExecuteTemplate(w, "trueHeader", data)
		_ = t.ExecuteTemplate(w, "userPageHead", pageOwner)
		if len(themes) != 0 {
			//fmt.Println("3 theme")
			_ = t.ExecuteTemplate(w, "userPageTheme", themes)
		} else {
			//fmt.Println("4 theme")
			_ = t.ExecuteTemplate(w, "emptyUserPageTheme", themes)
		}
		_ = t.ExecuteTemplate(w, "scripts", nil)
	}
}

func userPageAbout(w http.ResponseWriter, r *http.Request) {
	//session expiring update
	session, _ := store.Get(r, "session-name")
	fmt.Println("Index data:", session.Values["logged"], session.Values["username"])

	if session.Values["logged"] != true {
		//http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	session.Options = &sessions.Options{
		MaxAge: 60 * 60 * 2,
	}
	_ = session.Save(r, w)
	vars := mux.Vars(r)
	pageOwner, _, _ := getUserPage(vars["username"])

	t, _ := template.ParseFiles(
		"./web/templates/scripts.html",
		"./web/templates/trueHeader.html",
		"./web/templates/userPage.html",
		"./web/templates/userPageHeadAbout.html")
	data.Username = fmt.Sprint(session.Values["username"])
	if vars["username"] == fmt.Sprint(session.Values["username"]) {
		//fmt.Println("1 ab")
		_ = t.ExecuteTemplate(w, "trueHeader", data)
		_ = t.ExecuteTemplate(w, "userOwnPageHead", data)
		_ = t.ExecuteTemplate(w, "userPageAbout", pageOwner)
		_ = t.ExecuteTemplate(w, "scripts", nil)
	} else {
		//fmt.Println("2 ab")
		_ = t.ExecuteTemplate(w, "trueHeader", data)
		if pageOwner.Description.Valid {
			//fmt.Println("3 ab")
			_ = t.ExecuteTemplate(w, "userPageHead", pageOwner)
		} else {
			//fmt.Println("4 ab")
			_ = t.ExecuteTemplate(w, "userPageHead", pageOwner)
			_ = t.ExecuteTemplate(w, "emptyUserPageAbout", pageOwner)
		}
		_ = t.ExecuteTemplate(w, "scripts", nil)
	}
}

func userPageFollows(w http.ResponseWriter, r *http.Request) {
	//session expiring update
	session, _ := store.Get(r, "session-name")
	if session.Values["logged"] != true {
		//http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	session.Options = &sessions.Options{
		MaxAge: 60 * 60 * 2,
	}
	_ = session.Save(r, w)

	vars := mux.Vars(r)
	pageOwner, _, _ := getUserPage(vars["username"])

	t, _ := template.ParseFiles(
		"./web/templates/scripts.html",
		"./web/templates/trueHeader.html",
		"./web/templates/userPage.html",
		"./web/templates/userPageHeadFollows.html")
	data.Username = fmt.Sprint(session.Values["username"])
	if vars["username"] == fmt.Sprint(session.Values["username"]) {
		//fmt.Println("1 fol")
		_ = t.ExecuteTemplate(w, "trueHeader", data)
		_ = t.ExecuteTemplate(w, "userOwnPageHead", data)
		_ = t.ExecuteTemplate(w, "userPageFollows", pageOwner)
		_ = t.ExecuteTemplate(w, "scripts", nil)
	} else {
		//fmt.Println("2 fol")
		_ = t.ExecuteTemplate(w, "trueHeader", data)
		_ = t.ExecuteTemplate(w, "userPageHead", pageOwner)
		_ = t.ExecuteTemplate(w, "userPageFollows", pageOwner)
		_ = t.ExecuteTemplate(w, "scripts", nil)
	}
}

func userPageEdit(w http.ResponseWriter, r *http.Request) {
	//session expiring update
	session, _ := store.Get(r, "session-name")
	if session.Values["logged"] != true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	session.Options = &sessions.Options{
		MaxAge: 60 * 60 * 2,
	}
	_ = session.Save(r, w)
}

func themePage(w http.ResponseWriter, r *http.Request) {
	//session expiring update
	session, _ := store.Get(r, "session-name")
	fmt.Println("Theme page:", session.Values["logged"], session.Values["username"])

	if session.Values["logged"] != true {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	session.Options = &sessions.Options{
		MaxAge: 60 * 60 * 2,
	}
	_ = session.Save(r, w)

	//getting data from url
	vars := mux.Vars(r)
	//getting theme data from bd
	theme := getThemePage(vars["id"])

	t, _ := template.ParseFiles(
		"./web/templates/scripts.html",
		"./web/templates/trueHeader.html",
		"./web/templates/themePage.html")

	data.Username = fmt.Sprint(session.Values["username"])
	if theme.CreatorName == fmt.Sprint(session.Values["username"]) {
		//fmt.Println("my theme")
		_ = t.ExecuteTemplate(w, "trueHeader", data)
		_ = t.ExecuteTemplate(w, "themeOwnPage", theme)
		_ = t.ExecuteTemplate(w, "scripts", nil)
	} else {
		//fmt.Println("else's theme")
		_ = t.ExecuteTemplate(w, "trueHeader", data)
		_ = t.ExecuteTemplate(w, "themeOwnPage", theme)
		_ = t.ExecuteTemplate(w, "scripts", nil)
	}
}

func themeUpload(w http.ResponseWriter, r *http.Request) {
	session, e := store.Get(r, "session-name")
	if e != nil {
		fmt.Fprintln(w, "Session getting error")
	}
	t, _ := template.ParseFiles(
		"./web/templates/scripts.html",
		"./web/templates/trueHeader.html",
		"./web/templates/uploadPage.html")
	data.Username = fmt.Sprint(session.Values["username"])
	_ = t.ExecuteTemplate(w, "trueHeader", data)
	_ = t.ExecuteTemplate(w, "uploadPage", nil)
	_ = t.ExecuteTemplate(w, "scripts", nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	var themeData theme
	_ = r.ParseMultipartForm(64)
	themeData.Name = r.FormValue("name")
	themeData.Description.Valid = true
	themeData.Description.String = r.FormValue("description")

	session, e := store.Get(r, "session-name")
	themeData.CreatorName = fmt.Sprint(session.Values["username"])

	file, _, e := r.FormFile("wallpaper")
	if e != nil {
		fmt.Fprintln(w, e.Error())
	}
	defer file.Close()

	themeData.ID = getThemeId(themeData.CreatorName) + 1

	themeData.Path = fmt.Sprintf("./web/user_files/%s/theme_%s", themeData.CreatorName, strconv.Itoa(themeData.ID))
	e = os.MkdirAll(themeData.Path, 0777)
	if e != nil {
		panic(e)
	}

	var osFile *os.File
	osFile, e = ioutil.TempFile(themeData.Path, "*.jpg")

	fileBytes, e := ioutil.ReadAll(file)
	if e != nil {
		fmt.Fprintln(w, e.Error())
	}
	osFile.Write(fileBytes)
	osFile.Close()

	oldPath := strings.Replace(osFile.Name(), "\\", "/", -1)
	newPath := themeData.Path + "/wallpaper.jpg"
	e = os.Rename(oldPath, newPath)
	if e != nil {
		fmt.Println("Rename error", e.Error())
	}

	e = saveTheme(themeData)
	if e != nil {
		fmt.Fprintln(w, e.Error())
	} else {
		//correct part
		http.Redirect(w, r, "/theme/"+strconv.Itoa(themeData.ID), http.StatusSeeOther)
	}
}

func exit(w http.ResponseWriter, r *http.Request) {
	session, e := store.Get(r, "session-name")
	if e != nil {
		fmt.Fprintln(w, "Session getting error")
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	session.Save(r, w)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
