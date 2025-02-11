package services

import (
	"context"
	"encoding/json"
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

func (s *WhatsappService) SendOrderCustomerMessage(ctx context.Context, req *whatsapp.SendOrderCustomerMessageRequest) (*whatsapp.SendOrderCustomerMessageResponse, error) {
	fmt.Println("Whatsapp Service :  SendOrderCustomerMessage")
	fmt.Println("Whatsapp Service :  SendOrderCustomerMessage - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	accountSid := s.Config.Twilio_Account_SID
	authToken := s.Config.Twilio_Auth_Token

	// Initialize Twilio client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	// Format product list
	var productList string
	var total float64
	for _, p := range req.Products {
		productPrice := float64(p.Price) / 100
		productList += fmt.Sprintf("- %s: $%.2f x %d = $%.2f\n", p.Name, productPrice, p.Quantity, productPrice*float64(p.Quantity))
		total += productPrice * float64(p.Quantity)
	}

	// WhatsApp contact link
	contactLink := fmt.Sprintf("https://wa.me/%s", req.ContactNumber)

	// Create dynamic variables JSON
	contentVariables := map[string]string{
		"1": req.CustomerName,
		"2": productList,
		"3": fmt.Sprintf("$%.2f", total),
		"4": contactLink,
	}
	contentVariablesJSON, _ := json.Marshal(contentVariables)

	// Define message parameters
	params := &api.CreateMessageParams{}
	params.SetTo(fmt.Sprintf("whatsapp:+58%v", req.ToPhone))                 // Recipient
	params.SetFrom(fmt.Sprintf("whatsapp:%v", s.Config.Twilio_Phone_Number)) // Twilio WhatsApp Number
	params.SetContentSid("HXf5305de969b2c33adfe6138fac5781df")               // Content SID
	params.SetContentVariables(string(contentVariablesJSON))                 // Dynamic variables

	// Send WhatsApp message
	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Print message SID (Twilio's unique ID for the message)
	if resp.Sid != nil {
		fmt.Println("Message Sent! SID:", *resp.Sid)
	}

	fmt.Println("Whatsapp Service :  SendOrderCustomerMessage - SUCCESS")
	return &whatsapp.SendOrderCustomerMessageResponse{
		Status: http.StatusCreated,
	}, nil
}
