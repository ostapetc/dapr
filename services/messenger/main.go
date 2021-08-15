package main

import (
	"github.com/gorilla/mux"
	"github.com/ostapetc/dapr/messenger/handlers"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/messages", handlers.ListHandler).Methods("GET")
	router.HandleFunc("/messages", handlers.CreateHandler).Methods("POST")
	router.HandleFunc("/messages", handlers.DeleteAllHandler).Methods("DELETE")
	router.HandleFunc("/messages/fixtures", handlers.CreateFixturesHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
