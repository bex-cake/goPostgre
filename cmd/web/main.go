package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}


func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	PORT := ":8000"

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	router := http.NewServeMux()
	router.HandleFunc("/", app.home)
	router.HandleFunc("/posts", app.posts)
	router.HandleFunc("/posts/create", app.create)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := &http.Server{
		Addr: PORT,
		ErrorLog: errorLog,
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
