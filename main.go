package main

import (
	"log"
	"net/http"
	"relogio_mundial/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/localtime", handlers.HandleTimeForm).Methods("POST")
	log.Printf("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
