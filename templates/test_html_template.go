package main

import (
	"html/template"
	"os"
)

func main() {
	t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	t.ExecuteTemplate(os.Stdout, "T", "<i>alert('you have been pwned')</i>") // Hello, &lt;i&gt;alert(&#39;you have been pwned&#39;)&lt;/i&gt;!
}
