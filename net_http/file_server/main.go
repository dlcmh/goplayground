package main

import (
	"net/http"
)

// Port
const (
	Port      = ":8080"
	FileFoler = "var/www"
)

func fileServer() {
	http.FileServer(http.Dir(FileFoler))
}

func main() {
	http.ListenAndServe(Port, fileServer())
}
