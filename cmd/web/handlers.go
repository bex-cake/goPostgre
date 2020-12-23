package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.gohtml",
		"./ui/html/base.layout.gohtml",
		"./ui/html/footer.partial.gohtml",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func posts(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Show all posts"))
}

func create(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		w.Header().Set("Allowed methods", http.MethodPost)
		http.Error(w, "Method " + r.Method + " not allowed", http.StatusMethodNotAllowed)
		return
	}
	_, _ = w.Write([]byte("Create new post"))
}
