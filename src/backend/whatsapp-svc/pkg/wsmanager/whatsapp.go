package wsmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type WhatsappConversation struct {
	ID       string            `json:"id"`
	Messages []WhatsappMessage `json:"messages"`
}

type WhatsappMessage struct {
	Message    string    `json:"message"`
	Date       time.Time `json:"date"`
	WasReceipt bool      `json:"wasReceipt"`
	WasRead    bool      `json:"wasRead"`
	FromMe     bool      `json:"fromMe"`
}

func (c *Client) whatsappEventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	case *events.QR:
		if len(v.Codes) > 0 {
			c.SendServerMessage(v.Codes[0], QR)
		}
	// case *events.PairSuccess:
	// 	c.SendServerMessage("Conectado mano")
	// case *events.OfflineSyncCompleted:
	// 	contacts, _ := c.whatsappClient.Store.Contacts.GetAllContacts()
	// 	for k, v := range contacts {
	// 		c.SendServerMessage(fmt.Sprintf("key[%s] value[%s]\n", k, v.FullName), Message)
	// 	}
	// personMsg := map[string][]*events.Message
	// evt, err := c.whatsappClient.ParseWebMessage(chatJID, historyMsg.GetMessage())
	case *events.HistorySync:
		c.handleHistorySync(v)

		// if v.Data.Progress != nil {
		// 	c.SendServerMessage(fmt.Sprintf("Progreso = %v", *v.Data.Progress))
		// 	c.SendServerMessage("CONVERSTAIONS")
		// 	for _, conversation := range v.Data.Conversations {
		// 		c.SendServerMessage(fmt.Sprintf("%s = %s = %s", conversation.GetNewJID(), conversation.GetOldJID(), conversation.GetDisplayName()))
		// 		for _, msg := range conversation.Messages {
		// 			c.SendServerMessage(fmt.Sprintf("%s --- %s", msg.Message.GetMessage().GetConversation(), "hola"))
		// 		}
		// 	}
		//
		// 	c.SendServerMessage("V3 MESSAGES")
		// 	for _, conversation := range v.Data.StatusV3Messages {
		// 		c.SendServerMessage(fmt.Sprintf("%s = %s = %s", conversation.GetParticipant(), conversation.GetPushName(), conversation.GetMessage().GetConversation()))
		// 	}
		//
		// 	c.SendServerMessage("PUSH NAMES")
		// 	for _, pusha := range v.Data.Pushnames {
		// 		c.SendServerMessage(fmt.Sprintf("%s ////////////// %s", pusha.GetID(), pusha.GetPushname()))
		// 	}
		//
		// 	c.SendServerMessage("PROGRESS")
		// }
	default:
		var r = reflect.TypeOf(v)
		fmt.Println("----->")
		fmt.Printf("-----> EVENT TYPE = %v", r)
		fmt.Println("----->")
	}
}

func (c *Client) newWhatsappClient(container *sqlstore.Container) {
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(c.whatsappEventHandler)
	c.whatsappClient = client

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				// Render the QR code here
				// e.g. qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				// c.qrWhatsapp = evt.Code
				// qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				fmt.Println("----------------------------------")
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		fmt.Println("================================")
		fmt.Println("======= ALREADY CONNECTED ======")
		fmt.Println("================================")
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	// <-ch
}

func (c *Client) handleHistorySync(v *events.HistorySync) {
	if v.Data.Progress != nil {
		b, err := json.Marshal(v.Data.Conversations)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.SendServerMessage(string(b), Conversations)
	}
}
