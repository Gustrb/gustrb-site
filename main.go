package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var views = map[string]struct{}{}

func init() {
	// Get all the files in the templates folder
	files, err := os.ReadDir("templates")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		path := filepath.Join("templates", file.Name())
		views[path] = struct{}{}
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)

	log.Print("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	var view string
	if r.URL.Path == "/" {
		view = "index.html"
	} else {
		view = r.URL.Path + ".html"
	}

	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(view))

	if _, ok := views[fp]; !ok {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
