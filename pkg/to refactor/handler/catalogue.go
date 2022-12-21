package handler

//func Catalogue(w http.ResponseWriter, r *http.Request) {
//	//Session start
//	session, e := server.store.Get(r, "session-name")
//	error2.errorProc(w, e, "Session start error")
//
//	//Session expiring update
//	if auth.AuthCheck(session) {
//		auth.UpdateSession(session, w, r)
//	} else {
//		//Redirecting not auth users
//		http.Redirect(w, r, "/", http.StatusSeeOther)
//	}
//
//	lastThemes, e := database.GetLastThemes()
//	error2.errorProc(w, e, "Getting user page data error")
//
//	meowtThemes, e := database.GetCreatorsThemes("MEOWT")
//	error2.errorProc(w, e, "Getting user page data error")
//
//	//Parsing templates
//	t, e := template.ParseFiles(
//		"./web/templates/scripts.html",
//		"./web/templates/trueHeader.html",
//		"./web/templates/catalogue.html",
//		"./web/templates/catMeowt.html",
//		"./web/templates/catHead.html")
//	error2.errorProc(w, e, "Template parsing error")
//
//	//Executing templates with db data
//	var headerData struct {
//		Username string
//	}
//	headerData.Username = fmt.Sprint(session.Values["username"])
//
//	e = t.ExecuteTemplate(w, "trueHeader", headerData)
//	error2.errorProc(w, e, "Template executing error")
//
//	e = t.ExecuteTemplate(w, "catHead", nil)
//	error2.errorProc(w, e, "Template executing error")
//
//	e = t.ExecuteTemplate(w, "catalogue", lastThemes)
//	error2.errorProc(w, e, "Template executing error")
//
//	e = t.ExecuteTemplate(w, "catMeowt", meowtThemes)
//	error2.errorProc(w, e, "Template executing error")
//
//	e = t.ExecuteTemplate(w, "scripts", nil)
//	error2.errorProc(w, e, "Template executing error")
//}
