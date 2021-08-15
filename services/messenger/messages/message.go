package messages

type Message struct {
	Id          int    `json:"id"`
	SenderId    int    `json:"senderId" schema:"senderId,required"`
	RecipientId int    `json:"recipientId" schema:"recipientId,required"`
	Body        string `json:"body" schema:"body,required"`
}

func NewMessage(senderId int, recipientId int, body string) Message {
	return Message{
		SenderId:    senderId,
		RecipientId: recipientId,
		Body:        body,
	}
}
