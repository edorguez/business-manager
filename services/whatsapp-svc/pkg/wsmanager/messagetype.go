package wsmanager

type MessageType int

const (
	QR_CODE = iota
	CONVERSATIONS_CODE
	MESSAGE_CODE
	ALREADY_CONNECTED_CODE
)

var messageTypeName = map[MessageType]string{
	QR_CODE:                "QR_CODE",
	CONVERSATIONS_CODE:     "CONVERSATIONS_CODE",
	MESSAGE_CODE:           "MESSAGE_CODE",
	ALREADY_CONNECTED_CODE: "ALREADY_CONNECTED_CODE",
}

func (mt MessageType) String() string {
	return messageTypeName[mt]
}
