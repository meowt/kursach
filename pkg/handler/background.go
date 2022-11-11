package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func background(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseMultipartForm(64)
	file, _, e := r.FormFile("background")
	errorProc(w, e, "Forming file error")
	defer file.Close()

	session, e := store.Get(r, "session-name")
	OwnerName := fmt.Sprint(session.Values["username"])

	path := fmt.Sprintf("./web/user_files/%s", OwnerName)
	e = os.MkdirAll(path, 0777)
	if e != nil {
		panic(e)
	}

	var osFile *os.File
	osFile, e = os.CreateTemp(path, "*.jpg")
	errorProc(w, e, "Temping file error")

	fileBytes, e := io.ReadAll(file)
	errorProc(w, e, "Reading file error")

	_, e = osFile.Write(fileBytes)
	errorProc(w, e, "Writing file error")

	osFile.Close()

	oldPath := strings.Replace(osFile.Name(), "\\", "/", -1)
	newPath := path + "/background.jpg"
	e = os.Rename(oldPath, newPath)
	if e != nil {
		fmt.Println("Rename error", e.Error())
	}
	path = fmt.Sprintf("/user_files/%s", OwnerName)

	if e != nil {
		errorProc(w, e, "Saving theme error")
	} else {
		//correct part
		http.Redirect(w, r, "/user/"+OwnerName, http.StatusSeeOther)
	}
}
