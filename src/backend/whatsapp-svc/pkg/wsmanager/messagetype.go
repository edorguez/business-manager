package wsmanager

type MessageType int

const (
	QR = iota
	Conversations
	Message
)

var messageTypeName = map[MessageType]string{
	QR:            "QR",
	Conversations: "CONVERSATIONS",
	Message:       "MESSAGE",
}

func (mt MessageType) String() string {
	return messageTypeName[mt]
}
