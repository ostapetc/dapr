package main

import (
	"encoding/json"
	"fmt"
	"github.com/dapr/go-sdk/service/common"
	"github.com/gorilla/mux"
	"github.com/ostapetc/dapr/statistics/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dsstatus", handle).Methods("POST")

	log.Fatal(http.ListenAndServe(":8082", router))
}

func handle(w http.ResponseWriter, r *http.Request) {
	var event *common.TopicEvent

	err := json.NewDecoder(r.Body).Decode(&event)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("json decode error", err)
	}

	fmt.Println("Native handler")
	handler.HandleEvent(event)
}
