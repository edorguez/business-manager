package wsmanager

import (
	"context"
	"fmt"
	"reflect"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func (c *Client) whatsappEventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())
	case *events.QR:
		fmt.Println("dame qr")
		if len(v.Codes) > 0 {
			c.SendServerMessage(v.Codes[0])
		}
	case *events.PairSuccess:
		c.SendServerMessage("Conectado mano")
	case *events.OfflineSyncCompleted:
		contacts, _ := c.whatsappClient.Store.Contacts.GetAllContacts()
		for k, v := range contacts {
			c.SendServerMessage(fmt.Sprintf("key[%s] value[%s]\n", k, v.FullName))
		}
	// personMsg := map[string][]*events.Message
	// evt, err := c.whatsappClient.ParseWebMessage(chatJID, historyMsg.GetMessage())
	default:
		var r = reflect.TypeOf(v)
		fmt.Println("----->")
		fmt.Printf("-----> EVENT TYPE = %v\n", r)
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
