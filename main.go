package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Home page"))
}

func posts(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Show all posts"))
}

func create(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Method " + r.Method + " not allowed"))
		return
	}
	_, _ = w.Write([]byte("Create new post"))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", home)
	router.HandleFunc("/posts", posts)
	router.HandleFunc("/posts/create", create)

	server := &http.Server{
		Addr: ":8000",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
