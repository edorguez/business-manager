package wsmanager

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"time"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

type WhatsappConversation struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	ProfilePictureUrl string            `json:"profilePictureUrl"`
	UnreadCount       uint32            `json:"unreadCount"`
	Messages          []WhatsappMessage `json:"messages"`
}

type WhatsappMessage struct {
	ID         string    `jsong:"id"`
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
			c.SendServerMessage(v.Codes[0], QR_CODE)
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

		result := make([]WhatsappConversation, 0, len(v.Data.Conversations))

		for _, conversation := range v.Data.Conversations {
			var addConversation WhatsappConversation
			addConversation.ID = conversation.GetID()
			addConversation.UnreadCount = conversation.GetUnreadCount()

			jid, _ := types.ParseJID(addConversation.ID)

			re := regexp.MustCompile(`^\d+`)
			// Find the match
			idFormat := re.FindString(jid.String())
			addConversation.Name = idFormat

			if users, err := c.whatsappClient.GetUserInfo([]types.JID{jid}); err != nil {
				fmt.Println(err)
			} else {
				for _, user := range users {

					if picture_info, err := c.whatsappClient.GetProfilePictureInfo(jid, nil); err != nil || picture_info == nil {
						fmt.Println("error", jid, user.PictureID, picture_info, err)
					} else {
						addConversation.ProfilePictureUrl = picture_info.URL
					}
				}
			}

			addConversation.Messages = make([]WhatsappMessage, 0, len(conversation.Messages))

			for _, msg := range conversation.Messages {
				var wsmsg WhatsappMessage
				wsmsg.ID = msg.Message.GetKey().GetID()
				wsmsg.Date = time.Unix(int64(msg.Message.GetMessageTimestamp()), 0)
				wsmsg.FromMe = msg.Message.Key.GetFromMe()
				wsmsg.WasReceipt = msg.Message.GetUserReceipt()[0].ReceiptTimestamp != nil
				wsmsg.WasRead = msg.Message.GetUserReceipt()[0].ReadTimestamp != nil

				conversationMsg := msg.Message.GetMessage().GetConversation()
				conversationExtended := msg.Message.GetMessage().GetExtendedTextMessage().GetText()
				conversationEdited := msg.Message.GetMessage().GetEditedMessage().GetMessage().GetProtocolMessage().GetEditedMessage().GetExtendedTextMessage().GetText()

				if len(conversationMsg) > 0 {
					wsmsg.Message = conversationMsg
				} else if len(conversationExtended) > 0 {
					wsmsg.Message = conversationExtended
				} else {
					wsmsg.Message = conversationEdited
				}

				if len(wsmsg.Message) > 0 {
					addConversation.Messages = append([]WhatsappMessage{wsmsg}, addConversation.Messages...)
				}
			}

			result = append(result, addConversation)
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetEscapeHTML(false)

		if err := enc.Encode(result); err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		jsonStr := buf.String()
		jsonStr = jsonStr[:len(jsonStr)-1]

		c.SendServerMessage(jsonStr, CONVERSATIONS_CODE)
	}
}
