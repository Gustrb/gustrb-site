package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Name string
		}{
			Name: "Gustavo",
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
