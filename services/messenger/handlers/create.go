package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/ostapetc/dapr/messenger/dapr"
	"github.com/ostapetc/dapr/messenger/messages"
	"github.com/ostapetc/dapr/messenger/store"
	"net/http"
)

var decoder = schema.NewDecoder()

func writeResponse(w http.ResponseWriter, statusCode int, body []byte) {
	w.WriteHeader(statusCode)

	_, err := w.Write(body)
	if err != nil {
		panic(err)
	}
}

func getErrorBody(err error) []byte {
	body := make(map[string]string)
	body["error"] = err.Error()

	jsonBody, err := json.Marshal(body)

	if err != nil {
		panic(err)
	}

	return jsonBody
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	var message messages.Message

	err = decoder.Decode(&message, r.PostForm)

	if err != nil {
		writeResponse(w, http.StatusUnprocessableEntity, getErrorBody(err))
		return
	}

	store.AddMessage(message)

	eventData := []byte("\"{\\\"source\\\":\\\"library\\\"}\"")
	ctx := context.Background()

	err = dapr.GetClient().PublishEvent(ctx, "pubsub", "deathStarStatus", eventData);
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Published event")
	}

	res, err := http.Post("http://localhost:3500/v1.0/publish/pubsub/deathStarStatus", "application/json", bytes.NewBuffer([]byte("{\"source\":\"http client\"}")))
	if err != nil {
		fmt.Println("Post publish error", err)
	}

	fmt.Println("Post publish response", res.StatusCode)

	writeResponse(w, http.StatusCreated, []byte{})
}

