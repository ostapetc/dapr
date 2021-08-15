package handlers

import (
	"fmt"
	"github.com/ostapetc/dapr/messenger/store"
	"net/http"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	messages := store.GetRawMessages()

	_, err := fmt.Fprintf(w, string(messages))
	if err != nil {
		panic(err)
	}
}
