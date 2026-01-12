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
		CompanyID:         arg.CompanyID,
		Jid:               arg.JID,
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
		JID:                  c.Jid,
		Name:                 type_converter.NewString(c.Name),
		UnreadCount:          type_converter.NewInt32(c.UnreadCount),
		IsGroup:              type_converter.NewBool(c.IsGroup),
		ProfilePictureUrl:    type_converter.NewString(c.ProfilePictureUrl),
		LastMessageTimestamp: type_converter.NewTime(c.LastMessageTimestamp),
	}

	return res, nil
}

func (s *WhatsappMessagingService) CreateMessage(ctx context.Context, arg datatransfer.CreateMessageRequestDto) (int64, error) {
	fmt.Println("WhatsappMessaging Service :  CreateMessage")
	fmt.Println("WhatsappMessaging Service :  CreateMessage - Req")
	fmt.Println(arg)
	fmt.Println("----------------")

	params := db.CreateMessageParams{
		CompanyID:       arg.CompanyID,
		ConversationJid: arg.ConversationJID,
		RemoteJid:       arg.RemoteJID,
		FromMe:          type_converter.NewSqlNullBool(arg.FromMe),
		MessageText:     type_converter.NewSqlNullString(arg.MessageText),
		MediaUrl:        type_converter.NewSqlNullString(arg.MediaUrl),
		MediaCaption:    type_converter.NewSqlNullString(arg.MediaCaption),
		Timestamp:       int64(arg.Timestamp),
		ReceivedAt:      type_converter.NewSqlNullTime(arg.ReceivedAt),
		EditedAt:        type_converter.NewSqlNullTime(arg.EditedAt),
		IsForwarded:     type_converter.NewSqlNullBool(arg.IsForwarded),
		IsDeleted:       type_converter.NewSqlNullBool(arg.IsDeleted),
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

func (s *WhatsappMessagingService) GetMessagesByConversationJID(ctx context.Context, arg datatransfer.GetMessagesByConversationJIDRequestDto) ([]datatransfer.GetMessagesByConversationResponseDto, error) {
	fmt.Println("WhatsappMessaging Service :  GetmessagesbyconversationJID")
	fmt.Println("WhatsappMessaging Service :  GetMessagesByConversationJID - Req")
	fmt.Println(arg)
	fmt.Println("----------------")

	params := db.GetMessagesByConversationJIDParams{
		ConversationJid: arg.ConversationJID,
		Limit:           arg.Limit,
		Offset:          arg.Offset,
	}

	c, err := s.Repo.GetMessagesByConversationJID(ctx, params)
	if err != nil {
		fmt.Println("WhatsappMessaging Service :  GetMessagesByConversationJID - ERROR")
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("WhatsappMessaging Service :  GetMessagesByConversationJID - SUCCESS")

	messages := make([]datatransfer.GetMessagesByConversationResponseDto, 0, len(c))
	for _, v := range c {
		messages = append(messages, datatransfer.GetMessagesByConversationResponseDto{
			ID:              v.ID,
			CompanyID:       v.CompanyID,
			ConversationJID: v.ConversationJid,
			RemoteJid:       v.RemoteJid,
			FromMe:          type_converter.NewBool(v.FromMe),
			MessageText:     type_converter.NewString(v.MessageText),
			MediaUrl:        type_converter.NewString(v.MediaUrl),
			MediaCaption:    type_converter.NewString(v.MediaCaption),
			Timestamp:       uint64(v.Timestamp),
			ReceivedAt:      type_converter.NewTime(v.ReceivedAt),
			EditedAt:        type_converter.NewTime(v.EditedAt),
			IsForwarded:     type_converter.NewBool(v.IsForwarded),
			IsDeleted:       type_converter.NewBool(v.IsDeleted),
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
