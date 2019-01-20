package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Port
const (
	Port = ":8080"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/pages/{id:[0-9]+}", PageHandler)
	http.Handle("/", router)
	http.ListenAndServe(Port, nil)
}

// PageHandler test handles
func PageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Vars returns route variables for the current request, if any
	pageID := vars["id"]
	fileName := "files/" + pageID + ".html"
	fmt.Println(vars) // eg map[id:2] on visit to http://localhost:8080/pages/2
	http.ServeFile(w, r, fileName)
}
