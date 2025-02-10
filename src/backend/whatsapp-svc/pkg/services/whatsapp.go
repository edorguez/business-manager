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

	// Sample data (replace with your dynamic data)
	customerName := "John"
	products := []struct {
		Name     string
		Price    float64
		Quantity int
	}{
		{"Product A", 10.0, 2},
		{"Product B", 15.0, 1},
	}

	// Format product list
	var productList string
	var total float64
	for _, p := range products {
		productList += fmt.Sprintf("- %s: $%.2f x %d = $%.2f\n", p.Name, p.Price, p.Quantity, p.Price*float64(p.Quantity))
		total += p.Price * float64(p.Quantity)
	}

	// WhatsApp contact link
	contactNumber := "1234567890" // Replace with your contact number (E.164 format without "+")
	contactLink := fmt.Sprintf("https://wa.me/%s", contactNumber)

	// Create dynamic variables JSON
	contentVariables := map[string]string{
		"1": customerName,
		"2": productList,
		"3": fmt.Sprintf("$%.2f", total),
		"4": contactLink,
	}
	contentVariablesJSON, _ := json.Marshal(contentVariables)

	// Define message parameters
	params := &api.CreateMessageParams{}
	params.SetTo(fmt.Sprintf("whatsapp:+58%v", req.ToPhone))                 // Recipient
	params.SetFrom(fmt.Sprintf("whatsapp:%v", s.Config.Twilio_Phone_Number)) // Twilio WhatsApp Number
	params.SetContentSid("HX64a2489d722cfd55de289b7f8710a434")               // Content SID
	// params.SetBody(`{"1":"12/1","2":"3pm"}`)                                 // Dynamic variables
	params.SetContentVariables(string(contentVariablesJSON)) // Dynamic variables

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
