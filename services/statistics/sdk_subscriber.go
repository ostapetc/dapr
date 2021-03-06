package main

import (
	"context"
	"fmt"
	"github.com/ostapetc/dapr/statistics/handler"
	"log"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

// Subscription to tell the dapr what topic to subscribe.
// - PubsubName: is the name of the component configured in the metadata of pubsub.yaml.
// - Topic: is the name of the topic to subscribe.
// - Route: tell dapr where to request the API to publish the message to the subscriber when get a message from topic.
var sub = &common.Subscription{
	PubsubName: "pubsub",
	Topic:      "deathStarStatus",
	Route:      "/dsstatus",
}

func main() {
	s := daprd.NewService(":8082")

	if err := s.AddTopicEventHandler(sub, eventHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

func eventHandler(ctx context.Context, event *common.TopicEvent) (retry bool, err error) {
	fmt.Println("Native handler")
	handler.HandleEvent(event)
	return false, nil
}