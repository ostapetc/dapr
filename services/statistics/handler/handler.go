package handler

import (
	"github.com/dapr/go-sdk/service/common"
	"log"
)

func HandleEvent(event *common.TopicEvent) {
	log.Printf(
		"Event - Source: %s, PubsubName: %s, Topic: %s, ID: %s, Data: %s",
		event.Source,
		event.PubsubName,
		event.Topic,
		event.ID,
		event.Data,
	)
}