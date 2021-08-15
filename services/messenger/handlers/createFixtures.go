package handlers

import (
	"github.com/ostapetc/dapr/messenger/store"
	"net/http"
)

func CreateFixturesHandler(w http.ResponseWriter, r *http.Request) {
	store.CreateFixtures()
}