package handlers

import (
	"github.com/ostapetc/dapr/messenger/store"
	"net/http"
)

func DeleteAllHandler(w http.ResponseWriter, r *http.Request) {
	store.DeleteAllMessages()
}