package handler

import (
	error2 "Diploma/pkg/error"
	"Diploma/server"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func avatar(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseMultipartForm(64)
	file, _, e := r.FormFile("avatar")
	error2.errorProc(w, e, "Forming file error")
	defer file.Close()

	session, e := server.store.Get(r, "session-name")
	OwnerName := fmt.Sprint(session.Values["username"])

	path := fmt.Sprintf("./web/user_files/%s", OwnerName)
	e = os.MkdirAll(path, 0777)
	if e != nil {
		panic(e)
	}

	var osFile *os.File
	osFile, e = os.CreateTemp(path, "*.jpg")
	error2.errorProc(w, e, "Temping file error")

	fileBytes, e := io.ReadAll(file)
	error2.errorProc(w, e, "Reading file error")

	_, e = osFile.Write(fileBytes)
	error2.errorProc(w, e, "Writing file error")

	osFile.Close()

	oldPath := strings.Replace(osFile.Name(), "\\", "/", -1)
	newPath := path + "/avatar.jpg"
	e = os.Rename(oldPath, newPath)
	if e != nil {
		fmt.Println("Rename error", e.Error())
	}
	path = fmt.Sprintf("/user_files/%s", OwnerName)

	if e != nil {
		error2.errorProc(w, e, "Saving theme error")
	} else {
		//correct part
		http.Redirect(w, r, "/user/"+OwnerName, http.StatusSeeOther)
	}
}
