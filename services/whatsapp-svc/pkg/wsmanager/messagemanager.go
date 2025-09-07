package wsmanager

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

func (c *Client) SendServerMessage(msg string, msgType MessageType) {
	obj := NewMessageEvent{
		SendMessageEvent: SendMessageEvent{
			Message:     msg,
			From:        "server",
			MessageType: msgType.String(),
		},
		Sent: time.Now(),
	}
	result, _ := json.Marshal(obj)

	c.connection.WriteMessage(websocket.TextMessage, result)
}
