package main

import (
	"fmt"
	"net/http"
	"time"
)

// Port
const (
	Port = ":8080"
)

func serveDynamic(w http.ResponseWriter, r *http.Request) {
	response := "The time is now " + time.Now().String() // The time is now 2019-01-21 06:27:42.977964 +0800 +08 m=+10.007683740
	fmt.Fprintln(w, response)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func main() {
	http.HandleFunc("/static", serveStatic)
	http.HandleFunc("/", serveDynamic)
	http.ListenAndServe(Port, nil)
}
