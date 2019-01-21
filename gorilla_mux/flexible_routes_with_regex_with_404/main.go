package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Port
const (
	Port = ":8080"
)

// run with gin:
// gin -a 8080 -p 8081 run main.go
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/pages/{id:[0-9]+}", PageHandler)
	http.Handle("/", router)
	http.ListenAndServe(Port, nil)
}

// PageHandler test handles
func PageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fileName := "files/" + pageID + ".html"
	_, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err) // eg stat files/20.html: no such file or directory
		fileName = "files/404.html"
	}
	http.ServeFile(w, r, fileName)
}
