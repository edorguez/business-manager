package repository

import (
	"context"
	"time"

	"github.com/edorguez/business-manager/services/whatsapp-svc/pkg/datatransfer"
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

func (wr *WhatsappMessagingRepo) GetMessagesByConversationJID(ctx context.Context, arg db.GetMessagesByConversationJIDParams) ([]db.GetMessagesByConversationJIDRow, error) {
	var result []db.GetMessagesByConversationJIDRow

	err := wr.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetMessagesByConversationJID(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (wr *WhatsappMessagingRepo) BulkSaveConversations(ctx context.Context, params datatransfer.BulkConversationParamsDto) error {
	if len(params.Conversations) == 0 {
		return nil
	}

	// Prepare arrays for bulk insert
	companyIDs := make([]int64, 0, len(params.Conversations))
	jids := make([]string, 0, len(params.Conversations))
	names := make([]string, 0, len(params.Conversations))
	unreadCounts := make([]int32, 0, len(params.Conversations))
	isGroups := make([]bool, 0, len(params.Conversations))
	profilePictureURLs := make([]string, 0, len(params.Conversations))

	for _, conv := range params.Conversations {
		companyIDs = append(companyIDs, params.CompanyID)
		jids = append(jids, conv.JID)
		names = append(names, conv.Name)
		unreadCounts = append(unreadCounts, conv.UnreadCount)
		isGroups = append(isGroups, conv.IsGroup)
		profilePictureURLs = append(profilePictureURLs, conv.ProfilePictureURL)
	}

	bulkParams := db.BulkUpsertConversationsParams{
		CompanyIds:         companyIDs,
		Jids:               jids,
		Names:              names,
		UnreadCounts:       unreadCounts,
		IsGroups:           isGroups,
		ProfilePictureUrls: profilePictureURLs,
	}

	err := wr.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		return q.BulkUpsertConversations(ctx, bulkParams)
	})

	return err
}

func (wr *WhatsappMessagingRepo) BulkSaveMessages(ctx context.Context, params datatransfer.BulkMessageParamsDto) error {
	if len(params.Messages) == 0 {
		return nil
	}

	// Prepare arrays for bulk insert
	companyIDs := make([]int64, 0, len(params.Messages))
	conversationJids := make([]string, 0, len(params.Messages))
	remoteJIDs := make([]string, 0, len(params.Messages))
	fromMes := make([]bool, 0, len(params.Messages))
	messageTypes := make([]string, 0, len(params.Messages))
	messageTexts := make([]string, 0, len(params.Messages))
	mediaURLs := make([]string, 0, len(params.Messages))
	mediaCaptions := make([]string, 0, len(params.Messages))
	statuses := make([]string, 0, len(params.Messages))
	timestamps := make([]time.Time, 0, len(params.Messages))
	receivedAts := make([]time.Time, 0, len(params.Messages))
	editedAts := make([]time.Time, 0, len(params.Messages))
	isForwardeds := make([]bool, 0, len(params.Messages))
	isDeleteds := make([]bool, 0, len(params.Messages))

	for _, msg := range params.Messages {
		companyIDs = append(companyIDs, params.CompanyID)
		conversationJids = append(conversationJids, msg.ConversationJID)
		remoteJIDs = append(remoteJIDs, msg.RemoteJID)
		fromMes = append(fromMes, msg.FromMe)
		messageTypes = append(messageTypes, msg.MessageType)
		messageTexts = append(messageTexts, msg.MessageText)
		mediaURLs = append(mediaURLs, msg.MediaURL)
		mediaCaptions = append(mediaCaptions, msg.MediaCaption)
		statuses = append(statuses, msg.Status)
		timestamps = append(timestamps, msg.Timestamp)
		receivedAts = append(receivedAts, msg.ReceivedAt)
		editedAts = append(editedAts, msg.EditedAt)
		isForwardeds = append(isForwardeds, msg.IsForwarded)
		isDeleteds = append(isDeleteds, msg.IsDeleted)
	}

	bulkParams := db.BulkUpsertMessagesParams{
		CompanyIds:       companyIDs,
		ConversationJids: conversationJids,
		RemoteJids:       remoteJIDs,
		FromMes:          fromMes,
		MessageTypes:     messageTypes,
		MessageTexts:     messageTexts,
		MediaUrls:        mediaURLs,
		MediaCaptions:    mediaCaptions,
		Statuses:         statuses,
		Timestamps:       timestamps,
		ReceivedAts:      receivedAts,
		EditedAts:        editedAts,
		IsForwardeds:     isForwardeds,
		IsDeleteds:       isDeleteds,
	}

	err := wr.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		return q.BulkUpsertMessages(ctx, bulkParams)
	})

	return err
}

func (wr *WhatsappMessagingRepo) BulkSaveConversationsAndMessages(ctx context.Context, convParams datatransfer.BulkConversationParamsDto, msgParams datatransfer.BulkMessageParamsDto) error {
	return wr.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		// Save conversations first
		if len(convParams.Conversations) > 0 {
			// Prepare conversation arrays
			companyIDs := make([]int64, 0, len(convParams.Conversations))
			jids := make([]string, 0, len(convParams.Conversations))
			names := make([]string, 0, len(convParams.Conversations))
			unreadCounts := make([]int32, 0, len(convParams.Conversations))
			isGroups := make([]bool, 0, len(convParams.Conversations))
			profilePictureURLs := make([]string, 0, len(convParams.Conversations))

			for _, conv := range convParams.Conversations {
				companyIDs = append(companyIDs, convParams.CompanyID)
				jids = append(jids, conv.JID)
				names = append(names, conv.Name)
				unreadCounts = append(unreadCounts, conv.UnreadCount)
				isGroups = append(isGroups, conv.IsGroup)
				profilePictureURLs = append(profilePictureURLs, conv.ProfilePictureURL)
			}

			convBulkParams := db.BulkUpsertConversationsParams{
				CompanyIds:         companyIDs,
				Jids:               jids,
				Names:              names,
				UnreadCounts:       unreadCounts,
				IsGroups:           isGroups,
				ProfilePictureUrls: profilePictureURLs,
			}

			if err := q.BulkUpsertConversations(ctx, convBulkParams); err != nil {
				return err
			}
		}

		// Save messages
		if len(msgParams.Messages) > 0 {
			// Prepare message arrays
			companyIDs := make([]int64, 0, len(msgParams.Messages))
			conversationJids := make([]string, 0, len(msgParams.Messages))
			remoteJIDs := make([]string, 0, len(msgParams.Messages))
			fromMes := make([]bool, 0, len(msgParams.Messages))
			messageTypes := make([]string, 0, len(msgParams.Messages))
			messageTexts := make([]string, 0, len(msgParams.Messages))
			mediaURLs := make([]string, 0, len(msgParams.Messages))
			mediaCaptions := make([]string, 0, len(msgParams.Messages))
			statuses := make([]string, 0, len(msgParams.Messages))
			timestamps := make([]time.Time, 0, len(msgParams.Messages))
			receivedAts := make([]time.Time, 0, len(msgParams.Messages))
			editedAts := make([]time.Time, 0, len(msgParams.Messages))
			isForwardeds := make([]bool, 0, len(msgParams.Messages))
			isDeleteds := make([]bool, 0, len(msgParams.Messages))

			for _, msg := range msgParams.Messages {
				companyIDs = append(companyIDs, msgParams.CompanyID)
				conversationJids = append(conversationJids, msg.ConversationJID)
				remoteJIDs = append(remoteJIDs, msg.RemoteJID)
				fromMes = append(fromMes, msg.FromMe)
				messageTypes = append(messageTypes, msg.MessageType)
				messageTexts = append(messageTexts, msg.MessageText)
				mediaURLs = append(mediaURLs, msg.MediaURL)
				mediaCaptions = append(mediaCaptions, msg.MediaCaption)
				statuses = append(statuses, msg.Status)
				timestamps = append(timestamps, msg.Timestamp)
				receivedAts = append(receivedAts, msg.ReceivedAt)
				editedAts = append(editedAts, msg.EditedAt)
				isForwardeds = append(isForwardeds, msg.IsForwarded)
				isDeleteds = append(isDeleteds, msg.IsDeleted)
			}

			msgBulkParams := db.BulkUpsertMessagesParams{
				CompanyIds:       companyIDs,
				ConversationJids: conversationJids,
				RemoteJids:       remoteJIDs,
				FromMes:          fromMes,
				MessageTypes:     messageTypes,
				MessageTexts:     messageTexts,
				MediaUrls:        mediaURLs,
				MediaCaptions:    mediaCaptions,
				Statuses:         statuses,
				Timestamps:       timestamps,
				ReceivedAts:      receivedAts,
				EditedAts:        editedAts,
				IsForwardeds:     isForwardeds,
				IsDeleteds:       isDeleteds,
			}

			if err := q.BulkUpsertMessages(ctx, msgBulkParams); err != nil {
				return err
			}
		}

		return nil
	})
}
