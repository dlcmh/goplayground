// https://gowebexamples.com/templates/

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Todo struct
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData struct
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

var data = TodoPageData{
	PageTitle: "My TODO list",
	Todos: []Todo{
		{Title: "<i>Task</i> 1", Done: false},
		{Title: "Task 2", Done: true},
		{Title: "Task 3", Done: true},
	},
}

func servePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html")
	// t := template.Must(template.ParseFiles("layout.html"))
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", servePage)

	fmt.Println("Listening on 8081")
	http.ListenAndServe(":8081", nil)
}
