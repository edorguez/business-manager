package repository

import (
	"context"

	db "github.com/edorguez/business-manager/services/whatsapp-svc/pkg/db/sqlc"
)

type WhatsappMessagingRepo struct {
	SQLStorage *db.SQLStorage
}

func NewWhatsappMessagingRepo(sql *db.SQLStorage) *WhatsappMessagingRepo {
	return &WhatsappMessagingRepo{
		SQLStorage: sql,
	}
}

func (wr *WhatsappMessagingRepo) CreateConversation(ctx context.Context, arg db.CreateConversationParams) (int64, error) {
	var result int64

	err := wr.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.CreateConversation(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (wr *WhatsappMessagingRepo) GetConversationByJID(ctx context.Context, arg db.GetConversationByJIDParams) (db.GetConversationByJIDRow, error) {
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

func (wr *WhatsappMessagingRepo) GetConversationsByUser(ctx context.Context, arg db.GetConversationsByUserParams) ([]db.GetConversationsByUserRow, error) {
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

func (wr *WhatsappMessagingRepo) CreateMessage(ctx context.Context, arg db.CreateMessageParams) (int64, error) {
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

func (wr *WhatsappMessagingRepo) GetMessagesByConversation(ctx context.Context, arg db.GetMessagesByConversationParams) ([]db.GetMessagesByConversationRow, error) {
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
