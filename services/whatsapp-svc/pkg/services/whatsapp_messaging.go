package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

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

func (s *WhatsappMessagingService) GetConversationByJID(ctx context.Context, arg datatransfer.GetConversationByJIDRequestDto) (datatransfer.GetConversationByJIDResponseDto, error) {
	var result db.GetConversationByJIDRow

	err := wr.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetConversationByJID(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (s *WhatsappMessagingService) GetConversationsByUser(ctx context.Context, arg datatransfer.GetConversationsByUserRequestDto) ([]datatransfer.GetConversationsByUserResponseDto, error) {
	var result []db.GetConversationsByUserRow

	err := wr.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetConversationsByUser(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (s *WhatsappMessagingService) CreateMessage(ctx context.Context, arg datatransfer.CreateMessageRequestDto) (int64, error) {
	var result int64

	err := wr.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.CreateMessage(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (s *WhatsappMessagingService) GetMessagesByConversation(ctx context.Context, arg datatransfer.GetMessagesByConversationRequestDto) ([]datatransfer.GetMessagesByConversationResponseDto, error) {
	var result []db.GetMessagesByConversationRow

	err := wr.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetMessagesByConversation(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
