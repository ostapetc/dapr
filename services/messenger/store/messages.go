package store

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ostapetc/dapr/messenger/dapr"
	"github.com/ostapetc/dapr/messenger/messages"
)

const StoreName = "statestore"
const StoreKeyMessages = "messages"

var currId = 0

func GetRawMessages() []byte {
	item, err := dapr.GetClient().GetState(context.Background(), StoreName, StoreKeyMessages)
	if err != nil {
		panic(err)
	}

	return item.Value
}

func GetMessages() []*messages.Message {
	var msgs []*messages.Message

	_ = json.Unmarshal(GetRawMessages(), &msgs)

	return msgs
}

func AddMessage(message messages.Message) {
	currId++

	message.Id = currId

	msgs := GetMessages()
	msgs = append(msgs, &message)

	msgsJson, err := json.Marshal(msgs)
	if err != nil {
		panic(err)
	}

	err = dapr.GetClient().SaveState(context.Background(), StoreName, StoreKeyMessages, msgsJson)
	if err != nil {
		panic(err)
	}
}

func DeleteAllMessages() {
	err := dapr.GetClient().DeleteState(context.Background(), StoreName, StoreKeyMessages)
	if err != nil {
		panic(err)
	}
}

func CreateFixtures() {
	DeleteAllMessages()

	for i := 1; i <= 10; i++ {
		AddMessage(
			messages.NewMessage(
				i+100,
				i+200,
				fmt.Sprintf("Body %d", i),
			),
		)
	}
}
