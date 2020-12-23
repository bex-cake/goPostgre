package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", home)
	router.HandleFunc("/posts", posts)
	router.HandleFunc("/posts/create", create)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))



	server := &http.Server{
		Addr: ":8000",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
