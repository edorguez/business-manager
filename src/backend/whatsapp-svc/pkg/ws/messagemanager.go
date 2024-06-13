package ws

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

func (c *Client) SendServerMessage(msg string) {
	obj := NewMessageEvent{
		SendMessageEvent: SendMessageEvent{
			Message: msg,
			From:    "server",
		},
		Sent: time.Now(),
	}
	result, _ := json.Marshal(obj)

	c.connection.WriteMessage(websocket.TextMessage, result)
}
