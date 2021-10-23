package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

//go:embed templates
var templates embed.FS

func GetTemplates() (fs.FS, error) {
	f, err := fs.Sub(templates, "templates")

	if err != nil {
		return nil, err
	}

	return f, nil
}

func handleIndex(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "Not Found", 404)
			return
		}

		w.Header().Add("content-type", "text/html; charset=utf-8")
		t.ExecuteTemplate(w, "home.html.tmpl", nil)
	}
}

func handleAbout(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/html; charset=utf-8")
		t.ExecuteTemplate(w, "about.html.tmpl", nil)
	}
}

func main() {
	tfs, _ := GetTemplates()
	t, err := template.New("").ParseFS(tfs, "*.html.tmpl")

	log.Println(t.DefinedTemplates())

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/about", handleAbout(t))
	mux.HandleFunc("/", handleIndex(t))

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
