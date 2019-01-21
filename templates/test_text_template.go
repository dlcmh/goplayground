package main

import (
	"os"
	"text/template"
)

func main() {
	t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	t.ExecuteTemplate(os.Stdout, "T", "<i>alert('you have been pwned')</i>") // Hello, <i>alert('you have been pwned')</i>!
}
