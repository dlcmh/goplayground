package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbconn = "dbname=goplayground sslmode=disable" // sslmode=disable, otherwise, pq: SSL is not enabled on the server
	port   = ":8080"
)

type page struct {
	Title       string
	Content     string
	PublishedOn string
}

func connectToDatabase() {
	database, err := sql.Open("postgres", dbconn)
	if err != nil {
		log.Println("Couldn't connect!")
		log.Fatal(err.Error)
	}
	db = database
}

func servePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := page{}
	err :=
		db.QueryRow("SELECT title, content, published_on FROM pages WHERE guid=$1", pageGUID).
			Scan(&thisPage.Title, &thisPage.Content, &thisPage.PublishedOn)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Couldn't fetch page:", pageGUID)
		return
	}
	html := `<html><head><title>` + thisPage.Title + `</title></head><body><h1>` + thisPage.Title +
		`</h1><div>` + thisPage.Content + `</div></body></html>`
	fmt.Fprintln(w, html)
}

func main() {
	connectToDatabase()
	routes := mux.NewRouter()
	routes.HandleFunc("/page/{guid:[0-9a-zA\\-]+}", servePage)
	http.Handle("/", routes)
	http.ListenAndServe(port, nil)
}
