package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/EdoRguez/business-manager/whatsapp-svc/pkg/config"
	"github.com/EdoRguez/business-manager/whatsapp-svc/pkg/pb/whatsapp"
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

	// Replace with your actual token and phone number ID
	token := s.Config.Token
	phoneNumberID := s.Config.Phone_Id
	url := fmt.Sprintf("https://graph.facebook.com/v21.0/%s/messages", phoneNumberID)

	// JSON payload
	payload := []byte(`{
		"messaging_product": "whatsapp",
		"to": "584141280555",
		"type": "template",
		"template": {
			"name": "hello_world",
			"language": {
				"code": "en_US"
			}
		}
	}`)

	// Create a new HTTP request
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return &whatsapp.SendMessageResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	// Set headers
	httpReq.Header.Set("Authorization", "Bearer "+token)
	httpReq.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println("Error sending request:", err)
		// return
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Println("Response Status:", resp.Status)

	// Optionally, read and print the response body
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))

	fmt.Println("Whatsapp Service :  SendMessage - SUCCESS")
	return &whatsapp.SendMessageResponse{
		Status: http.StatusCreated,
	}, nil
}
