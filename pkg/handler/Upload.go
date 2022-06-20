package handler

import (
	"Kursach/pkg/database"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	var themeData database.Theme
	_ = r.ParseMultipartForm(64)
	themeData.Name = r.FormValue("name")
	themeData.Description.Valid = true
	themeData.Description.String = r.FormValue("description")

	session, e := store.Get(r, "session-name")
	themeData.CreatorName = fmt.Sprint(session.Values["username"])

	file, _, e := r.FormFile("wallpaper")
	errorProc(w, e, "Forming file error")
	defer file.Close()

	themeData.ID = database.GetLastThemeId() + 1

	themeData.Path = fmt.Sprintf("./web/user_files/%s/theme_%s", themeData.CreatorName, strconv.Itoa(themeData.ID))
	e = os.MkdirAll(themeData.Path, 0777)
	if e != nil {
		panic(e)
	}

	var osFile *os.File
	osFile, e = ioutil.TempFile(themeData.Path, "*.jpg")
	errorProc(w, e, "Temping file error")

	fileBytes, e := ioutil.ReadAll(file)
	errorProc(w, e, "Reading file error")

	_, e = osFile.Write(fileBytes)
	errorProc(w, e, "Writing file error")

	osFile.Close()

	oldPath := strings.Replace(osFile.Name(), "\\", "/", -1)
	newPath := themeData.Path + "/wallpaper.jpg"
	e = os.Rename(oldPath, newPath)
	if e != nil {
		fmt.Println("Rename error", e.Error())
	}
	themeData.Path = fmt.Sprintf("/user_files/%s/theme_%s/", themeData.CreatorName, strconv.Itoa(themeData.ID))

	themeData.ReleaseDate.Time = time.Now()
	e = database.SaveTheme(themeData)
	if e != nil {
		errorProc(w, e, "Saving theme error")
	} else {
		//correct part
		http.Redirect(w, r, "/theme/"+strconv.Itoa(themeData.ID), http.StatusSeeOther)
	}
}
