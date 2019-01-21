package main

import (
	"database/sql"
	"html/template"
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
	RawContent  string
	Content     template.HTML
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
			Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.PublishedOn)
	thisPage.Content = template.HTML(thisPage.RawContent)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Couldn't fetch page:", pageGUID)
		return
	}
	t, _ := template.ParseFiles("templates/blog.html")
	t.Execute(w, thisPage)
}

func main() {
	connectToDatabase()
	routes := mux.NewRouter()
	routes.HandleFunc("/page/{guid:[0-9a-zA\\-]+}", servePage)
	http.Handle("/", routes)
	http.ListenAndServe(port, nil)
}
