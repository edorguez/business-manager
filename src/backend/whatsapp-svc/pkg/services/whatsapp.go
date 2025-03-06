package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/EdoRguez/business-manager/whatsapp-svc/pkg/config"
	db "github.com/EdoRguez/business-manager/whatsapp-svc/pkg/db/sqlc"
	"github.com/EdoRguez/business-manager/whatsapp-svc/pkg/pb/whatsapp"
	repo "github.com/EdoRguez/business-manager/whatsapp-svc/pkg/repository"
	"github.com/EdoRguez/business-manager/whatsapp-svc/pkg/util/phone"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type WhatsappService struct {
	Config *config.Config
	Repo   *repo.WhatsappRepo
	whatsapp.UnimplementedWhatsappServiceServer
}

func (s *WhatsappService) SendOrderCustomerMessage(ctx context.Context, req *whatsapp.SendOrderCustomerMessageRequest) (*whatsapp.SendOrderCustomerMessageResponse, error) {
	fmt.Println("Whatsapp Service :  SendOrderCustomerMessage")
	fmt.Println("Whatsapp Service :  SendOrderCustomerMessage - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	accountSid := s.Config.Twilio_Account_SID
	authToken := s.Config.Twilio_Auth_Token

	c, err := s.Repo.GetBusinessPhoneByCompanyId(ctx, req.CompanyId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Whatsapp Service :  GetBusinessPhoneByCompanyId - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Company number not found"
		}

		return &whatsapp.SendOrderCustomerMessageResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

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
		productList += fmt.Sprintf("- %s: $%.2f x %d = $%.2f\r", p.Name, productPrice, p.Quantity, productPrice*float64(p.Quantity))
		total += productPrice * float64(p.Quantity)
	}

	// Create dynamic variables JSON
	contentVariables := map[string]string{
		"1": req.CustomerName,
		"2": productList,
		"3": fmt.Sprintf("$%.2f", total),
		"4": c.Phone,
	}
	contentVariablesJSON, _ := json.Marshal(contentVariables)

	// Define message parameters
	params := &api.CreateMessageParams{}
	params.SetTo(fmt.Sprintf("whatsapp:+58%v", phone.RemovePhoneZero(req.ToPhone))) // Recipient
	params.SetFrom(fmt.Sprintf("whatsapp:%v", s.Config.Twilio_Phone_Number))        // Twilio WhatsApp Number
	params.SetContentSid("HX3929fc4860efca29f7312839fa3a0827")                      // Content SID
	params.SetContentVariables(string(contentVariablesJSON))                        // Dynamic variables

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

func (s *WhatsappService) SendOrderBusinessMessage(ctx context.Context, req *whatsapp.SendOrderBusinessMessageRequest) (*whatsapp.SendOrderBusinessMessageResponse, error) {
	fmt.Println("Whatsapp Service :  SendOrderBusinessMessage")
	fmt.Println("Whatsapp Service :  SendOrderBusinessMessage - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	accountSid := s.Config.Twilio_Account_SID
	authToken := s.Config.Twilio_Auth_Token

	c, err := s.Repo.GetBusinessPhoneByCompanyId(ctx, req.CompanyId)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Whatsapp Service :  GetBusinessPhoneByCompanyId - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Company number not found"
		}

		return &whatsapp.SendOrderBusinessMessageResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

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
		productList += fmt.Sprintf("- %s: $%.2f x %d = $%.2f\r", p.Name, productPrice, p.Quantity, productPrice*float64(p.Quantity))
		total += productPrice * float64(p.Quantity)
	}

	// Create dynamic variables JSON
	contentVariables := map[string]string{
		"1": req.CustomerName,
		"2": productList,
		"3": fmt.Sprintf("$%.2f", total),
		"4": req.ContactNumber,
	}
	contentVariablesJSON, _ := json.Marshal(contentVariables)

	// Define message parameters
	params := &api.CreateMessageParams{}
	params.SetTo(fmt.Sprintf("whatsapp:+58%v", phone.RemovePhoneZero(c.Phone))) // Recipient
	params.SetFrom(fmt.Sprintf("whatsapp:%v", s.Config.Twilio_Phone_Number))    // Twilio WhatsApp Number
	params.SetContentSid("HX15a23f804b9a05bb807a843c09ac5975")                  // Content SID
	params.SetContentVariables(string(contentVariablesJSON))                    // Dynamic variables

	// Send WhatsApp message
	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Print message SID (Twilio's unique ID for the message)
	if resp.Sid != nil {
		fmt.Println("Message Sent! SID:", *resp.Sid)
	}

	fmt.Println("Whatsapp Service :  SendOrderBusinessMessage - SUCCESS")
	return &whatsapp.SendOrderBusinessMessageResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *WhatsappService) CreateBusinessPhone(ctx context.Context, req *whatsapp.CreateBusinessPhoneRequest) (*whatsapp.CreateBusinessPhoneResponse, error) {
	fmt.Println("Whatsapp Service :  CreateBusinessPhone")
	fmt.Println("Whatsapp Service :  CreateBusinessPhone - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	existBusiness, errExist := s.Repo.GetBusinessPhoneByCompanyId(ctx, req.CompanyId)

	if errExist != nil && errExist != sql.ErrNoRows {
		fmt.Println("Whatsapp Service :  CreateBusinessPhone - ERROR")
		fmt.Println(errExist.Error())
		return &whatsapp.CreateBusinessPhoneResponse{
			Status: http.StatusConflict,
			Error:  errExist.Error(),
		}, nil
	}

	if existBusiness.ID != 0 {
		fmt.Println("Whatsapp Service :  CreateBusinessPhone - ERROR")
		fmt.Println("Business already exists")
		return &whatsapp.CreateBusinessPhoneResponse{
			Status: http.StatusConflict,
			Error:  "Business already exists",
		}, nil
	}

	createBusinessPhoneParams := db.CreateBusinessPhoneParams{
		CompanyID: req.CompanyId,
		Phone:     req.Phone,
	}

	_, err := s.Repo.CreateBusinessPhone(ctx, createBusinessPhoneParams)
	if err != nil {
		fmt.Println("Whatsapp Service :  CreateBusinessPhone - ERROR")
		fmt.Println(err.Error())
		return &whatsapp.CreateBusinessPhoneResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Whatsapp Service :  CreateBusinessPhone - SUCCESS")
	return &whatsapp.CreateBusinessPhoneResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *WhatsappService) GetBusinessPhoneByCompanyId(ctx context.Context, req *whatsapp.GetBusinessPhoneByCompanyIdRequest) (*whatsapp.GetBusinessPhoneByCompanyIdResponse, error) {
	fmt.Println("Whatsapp Service :  GetBusinessPhoneByCompanyId")
	fmt.Println("Whatsapp Service :  GetBusinessPhoneByCompanyId - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	c, err := s.Repo.GetBusinessPhoneByCompanyId(ctx, req.CompanyId)
	if err != nil {
		fmt.Println("Whatsapp Service :  GetBusinessPhoneByCompanyId - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &whatsapp.GetBusinessPhoneByCompanyIdResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Whatsapp Service :  GetBusinessPhoneByCompanyId - SUCCESS")
	return &whatsapp.GetBusinessPhoneByCompanyIdResponse{
		Id:        c.ID,
		CompanyId: c.CompanyID,
		Phone:     c.Phone,
		Status:    http.StatusOK,
	}, nil
}

func (s *WhatsappService) UpdateBusinessPhone(ctx context.Context, req *whatsapp.UpdateBusinessPhoneRequest) (*whatsapp.UpdateBusinessPhoneResponse, error) {
	fmt.Println("Whatsapp Service :  UpdateBusinessPhone")
	fmt.Println("Whatsapp Service :  UpdateBusinessPhone - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.UpdateBusinessPhoneParams{
		CompanyID: req.CompanyId,
		Phone:     req.Phone,
	}

	_, err := s.Repo.UpdateBusinessPhone(ctx, params)
	if err != nil {
		fmt.Println("Whatsapp Service :  UpdateBusinessPhone - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}
		return &whatsapp.UpdateBusinessPhoneResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Whatsapp Service :  UpdateBusinessPhone - SUCCESS")
	return &whatsapp.UpdateBusinessPhoneResponse{
		Status: http.StatusNoContent,
	}, nil
}
