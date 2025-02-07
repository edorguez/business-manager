package services

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/EdoRguez/business-manager/whatsapp-svc/pkg/config"
	"github.com/EdoRguez/business-manager/whatsapp-svc/pkg/pb/whatsapp"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type WhatsappService struct {
	Config *config.Config
	whatsapp.UnimplementedWhatsappServiceServer
}

func (s *WhatsappService) SendMessage(ctx context.Context, req *whatsapp.SendMessageRequest) (*whatsapp.SendMessageResponse, error) {
	fmt.Println("Whatsapp Service :  SendMessage")
	fmt.Println("Whatsapp Service :  SendMessage - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	accountSid := s.Config.Twilio_Account_SID
	authToken := s.Config.Twilio_Auth_Token

	// Initialize Twilio client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	// Define message parameters
	params := &api.CreateMessageParams{}
	params.SetTo("whatsapp:+584141280555")                                   // Recipient
	params.SetFrom(fmt.Sprintf("whatsapp:%v", s.Config.Twilio_Phone_Number)) // Twilio WhatsApp Number
	params.SetContentSid("HX350d429d32e64a552466cafecbe95f3c")               // Content SID
	params.SetBody(`{"1":"12/1","2":"3pm"}`)                                 // Dynamic variables

	// Send WhatsApp message
	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Print message SID (Twilio's unique ID for the message)
	if resp.Sid != nil {
		fmt.Println("Message Sent! SID:", *resp.Sid)
	}

	fmt.Println("Whatsapp Service :  SendMessage - SUCCESS")
	return &whatsapp.SendMessageResponse{
		Status: http.StatusCreated,
	}, nil
}
