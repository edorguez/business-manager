package services

import (
	"context"
	"fmt"

	"github.com/edorguez/business-manager/services/whatsapp-svc/pkg/datatransfer"
	db "github.com/edorguez/business-manager/services/whatsapp-svc/pkg/db/sqlc"
	repo "github.com/edorguez/business-manager/services/whatsapp-svc/pkg/repository"
	"github.com/edorguez/business-manager/shared/util/type_converter"
)

type WhatsappMessagingService struct {
	Repo *repo.WhatsappMessagingRepo
}

func (s *WhatsappMessagingService) CreateConversation(ctx context.Context, arg datatransfer.CreateConversationRequestDto) (int64, error) {
	fmt.Println("WhatsappMessaging Service :  CreateConversation")
	fmt.Println("WhatsappMessaging Service :  CreateConversation - Req")
	fmt.Println(arg)
	fmt.Println("----------------")

	params := db.CreateConversationParams{
		ID:                arg.ID,
		CompanyID:         arg.CompanyID,
		UserID:            arg.UserID,
		Jid:               arg.Jid,
		Name:              type_converter.NewSqlNullString(arg.Name),
		UnreadCount:       type_converter.NewSqlNullInt32(arg.UnreadCount),
		IsGroup:           type_converter.NewSqlNullBool(arg.IsGroup),
		ProfilePictureUrl: type_converter.NewSqlNullString(arg.ProfilePictureUrl),
	}

	id, err := s.Repo.CreateConversation(ctx, params)

	if err != nil {
		fmt.Println("WhatsappMessaging Service :  CreateConversation - ERROR")
		fmt.Println(err.Error())
		return id, err
	}

	fmt.Println("WhatsappMessaging Service :  CreateConversation - SUCCESS")
	return id, nil
}

func (s *WhatsappMessagingService) GetConversationByJID(ctx context.Context, arg datatransfer.GetConversationByJIDRequestDto) (*datatransfer.GetConversationByJIDResponseDto, error) {
	fmt.Println("WhatsappMessaging Service :  GetConversationByJID")
	fmt.Println("WhatsappMessaging Service :  GetConversationByJID - Req")
	fmt.Println(arg)
	fmt.Println("----------------")

	params := db.GetConversationByJIDParams{
		CompanyID: arg.CompanyID,
		UserID:    arg.UserID,
		Jid:       arg.Jid,
	}

	c, err := s.Repo.GetConversationByJID(ctx, params)
	if err != nil {
		fmt.Println("WhatsappMessaging Service :  GetConversationByJID - ERROR")
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("WhatsappMessaging Service :  GetConversationByJID - SUCCESS")

	res := &datatransfer.GetConversationByJIDResponseDto{
		ID:                   c.ID,
		CompanyID:            c.CompanyID,
		UserID:               c.UserID,
		Jid:                  c.Jid,
		Name:                 type_converter.NewString(c.Name),
		UnreadCount:          type_converter.NewInt32(c.UnreadCount),
		IsGroup:              type_converter.NewBool(c.IsGroup),
		ProfilePictureUrl:    type_converter.NewString(c.ProfilePictureUrl),
		LastMessageTimestamp: type_converter.NewTime(c.LastMessageTimestamp),
	}

	return res, nil
}

func (s *WhatsappMessagingService) Getconversationsbyuser(ctx context.Context, arg datatransfer.GetConversationsByUserRequestDto) ([]datatransfer.GetConversationsByUserResponseDto, error) {
	fmt.Println("WhatsappMessaging Service :  GetConversationByUser")
	fmt.Println("WhatsappMessaging Service :  GetConversationByUser - Req")
	fmt.Println(arg)
	fmt.Println("----------------")

	params := db.GetConversationsByUserParams{
		CompanyID: arg.CompanyID,
		UserID:    arg.UserID,
		Limit:     arg.Limit,
		Offset:    arg.Offset,
	}

	c, err := s.Repo.GetConversationsByUser(ctx, params)
	if err != nil {
		fmt.Println("WhatsappMessaging Service :  GetConversationByUser - ERROR")
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("WhatsappMessaging Service :  GetConversationByUser - SUCCESS")

	conversations := make([]datatransfer.GetConversationsByUserResponseDto, 0, len(c))
	for _, v := range c {
		conversations = append(conversations, datatransfer.GetConversationsByUserResponseDto{
			ID:                   v.ID,
			CompanyID:            v.CompanyID,
			UserID:               v.UserID,
			Jid:                  v.Jid,
			Name:                 type_converter.NewString(v.Name),
			UnreadCount:          type_converter.NewInt32(v.UnreadCount),
			IsGroup:              type_converter.NewBool(v.IsGroup),
			ProfilePictureUrl:    type_converter.NewString(v.ProfilePictureUrl),
			LastMessageTimestamp: type_converter.NewTime(v.LastMessageTimestamp),
		})
	}

	return conversations, nil
}

func (s *WhatsappMessagingService) CreateMessage(ctx context.Context, arg datatransfer.CreateMessageRequestDto) (int64, error) {
	fmt.Println("WhatsappMessaging Service :  CreateMessage")
	fmt.Println("WhatsappMessaging Service :  CreateMessage - Req")
	fmt.Println(arg)
	fmt.Println("----------------")

	params := db.CreateMessageParams{
		ID:             arg.ID,
		CompanyID:      arg.CompanyID,
		ConversationID: arg.ConversationID,
		MessageID:      arg.MessageID,
		RemoteJid:      arg.RemoteJid,
		FromMe:         type_converter.NewSqlNullBool(arg.FromMe),
		MessageType:    arg.MessageType,
		MessageText:    type_converter.NewSqlNullString(arg.MessageText),
		MediaUrl:       type_converter.NewSqlNullString(arg.MediaUrl),
		MediaCaption:   type_converter.NewSqlNullString(arg.MediaCaption),
		Status:         type_converter.NewSqlNullString(arg.Status),
		Timestamp:      arg.Timestamp,
		ReceivedAt:     type_converter.NewSqlNullTime(arg.ReceivedAt),
		EditedAt:       type_converter.NewSqlNullTime(arg.EditedAt),
		IsForwarded:    type_converter.NewSqlNullBool(arg.IsForwarded),
		IsDeleted:      type_converter.NewSqlNullBool(arg.IsDeleted),
	}

	id, err := s.Repo.CreateMessage(ctx, params)

	if err != nil {
		fmt.Println("WhatsappMessaging Service :  CreateMessage - ERROR")
		fmt.Println(err.Error())
		return id, err
	}

	fmt.Println("WhatsappMessaging Service :  CreateMessage - SUCCESS")
	return id, nil
}

func (s *WhatsappMessagingService) GetMessagesByConversation(ctx context.Context, arg datatransfer.GetMessagesByConversationRequestDto) ([]datatransfer.GetMessagesByConversationResponseDto, error) {
	fmt.Println("WhatsappMessaging Service :  Getmessagesbyconversation")
	fmt.Println("WhatsappMessaging Service :  GetMessagesByConversation - Req")
	fmt.Println(arg)
	fmt.Println("----------------")

	params := db.GetMessagesByConversationParams{
		ConversationID: arg.ConversationID,
		Limit:          arg.Limit,
		Offset:         arg.Offset,
	}

	c, err := s.Repo.GetMessagesByConversation(ctx, params)
	if err != nil {
		fmt.Println("WhatsappMessaging Service :  GetMessagesByConversation - ERROR")
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("WhatsappMessaging Service :  GetMessagesByConversation - SUCCESS")

	messages := make([]datatransfer.GetMessagesByConversationResponseDto, 0, len(c))
	for _, v := range c {
		messages = append(messages, datatransfer.GetMessagesByConversationResponseDto{
			ID:             v.ID,
			CompanyID:      v.CompanyID,
			ConversationID: v.ConversationID,
			MessageID:      v.MessageID,
			RemoteJid:      v.RemoteJid,
			FromMe:         type_converter.NewBool(v.FromMe),
			MessageType:    v.MessageType,
			MessageText:    type_converter.NewString(v.MessageText),
			MediaUrl:       type_converter.NewString(v.MediaUrl),
			MediaCaption:   type_converter.NewString(v.MediaCaption),
			Status:         type_converter.NewString(v.Status),
			Timestamp:      v.Timestamp,
			ReceivedAt:     type_converter.NewTime(v.ReceivedAt),
			EditedAt:       type_converter.NewTime(v.EditedAt),
			IsForwarded:    type_converter.NewBool(v.IsForwarded),
			IsDeleted:      type_converter.NewBool(v.IsDeleted),
		})
	}

	return messages, nil
}

func (s *WhatsappMessagingService) BulkSaveConversations(ctx context.Context, params datatransfer.BulkConversationParamsDto) error {
	return s.Repo.BulkSaveConversations(ctx, params)
}

func (s *WhatsappMessagingService) BulkSaveMessages(ctx context.Context, params datatransfer.BulkMessageParamsDto) error {
	return s.Repo.BulkSaveMessages(ctx, params)
}

func (s *WhatsappMessagingService) BulkSaveConversationsAndMessages(ctx context.Context, convParams datatransfer.BulkConversationParamsDto, msgParams datatransfer.BulkMessageParamsDto) error {
	return s.Repo.BulkSaveConversationsAndMessages(ctx, convParams, msgParams)
}
