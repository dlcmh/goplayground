package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbconn = "dbname=goplayground sslmode=disable" // sslmode=disable, otherwise, pq: SSL is not enabled on the server
)

type page struct {
	Title       string
	Content     string
	PublishedOn string
}

func fetchOneRow(guid string) {
	thisPage := page{}
	err :=
		db.QueryRow("SELECT title, content, published_on FROM pages WHERE guid=$1", guid).
			Scan(&thisPage.Title, &thisPage.Content, &thisPage.PublishedOn)
	if err != nil {
		log.Println("Couldn't fetch page:")
		log.Fatal(err.Error)
	}
	fmt.Println(thisPage.Title)
	fmt.Println(thisPage.Content)
	fmt.Println(thisPage.PublishedOn)
}

func main() {
	database, err := sql.Open("postgres", dbconn)
	if err != nil {
		log.Println("Couldn't connect!")
		log.Fatal(err.Error)
	}
	db = database
	fetchOneRow("hello-world")
}
