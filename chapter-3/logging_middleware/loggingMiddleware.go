package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func mainLogic(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing Request")
	w.Write([]byte("ok!"))
	log.Println("Finished Processing Request")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainLogic)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8080", loggedRouter)
}
