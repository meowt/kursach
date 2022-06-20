package handler

import (
	"fmt"
	"net/http"
)

func HandleAssets() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/"))))
	http.Handle("/user_files/", http.StripPrefix("/user_files/", http.FileServer(http.Dir("./web/user_files/"))))
} //Assets handler function

func errorProc(w http.ResponseWriter, e error, s string) {
	if e != nil {
		_, _ = fmt.Fprintln(w, e.Error())
		fmt.Println(s + "\n")
	}
} //Error processing function
