package datatransfer

import (
	"time"
)

type CreateConversationRequestDto struct {
	CompanyID         int64   `json:"company_id"`
	JID               string  `json:"jid"`
	Name              *string `json:"name"`
	UnreadCount       *int32  `json:"unread_count"`
	IsGroup           *bool   `json:"is_group"`
	ProfilePictureUrl *string `json:"profile_picture_url"`
}

type GetConversationByJIDRequestDto struct {
	ID        int64  `json:"id"`
	CompanyID int64  `json:"company_id"`
	Jid       string `json:"jid"`
}

type GetConversationByJIDResponseDto struct {
	ID                   int64      `json:"id"`
	CompanyID            int64      `json:"company_id"`
	JID                  string     `json:"jid"`
	Name                 *string    `json:"name"`
	UnreadCount          *int32     `json:"unread_count"`
	IsGroup              *bool      `json:"is_group"`
	ProfilePictureUrl    *string    `json:"profile_picture_url"`
	LastMessageTimestamp *time.Time `json:"last_message_timestamp"`
}

type CreateMessageRequestDto struct {
	CompanyID       int64      `json:"company_id"`
	ConversationJID string     `json:"conversation_jid"`
	RemoteJID       string     `json:"remote_jid"`
	FromMe          *bool      `json:"from_me"`
	MessageType     string     `json:"message_type"`
	MessageText     *string    `json:"message_text"`
	MediaUrl        *string    `json:"media_url"`
	MediaCaption    *string    `json:"media_caption"`
	Status          *string    `json:"status"`
	Timestamp       time.Time  `json:"timestamp"`
	ReceivedAt      *time.Time `json:"received_at"`
	EditedAt        *time.Time `json:"edited_at"`
	IsForwarded     *bool      `json:"is_forwarded"`
	IsDeleted       *bool      `json:"is_deleted"`
}

type GetMessagesByConversationJIDRequestDto struct {
	ConversationJID string `json:"conversation_jid"`
	Limit           int32  `json:"limit"`
	Offset          int32  `json:"offset"`
}

type GetMessagesByConversationResponseDto struct {
	ID              int64      `json:"id"`
	CompanyID       int64      `json:"company_id"`
	ConversationJID string     `json:"conversation_jid"`
	RemoteJid       string     `json:"remote_jid"`
	FromMe          *bool      `json:"from_me"`
	MessageType     string     `json:"message_type"`
	MessageText     *string    `json:"message_text"`
	MediaUrl        *string    `json:"media_url"`
	MediaCaption    *string    `json:"media_caption"`
	Status          *string    `json:"status"`
	Timestamp       time.Time  `json:"timestamp"`
	ReceivedAt      *time.Time `json:"received_at"`
	EditedAt        *time.Time `json:"edited_at"`
	IsForwarded     *bool      `json:"is_forwarded"`
	IsDeleted       *bool      `json:"is_deleted"`
}

type BulkConversationParamsDto struct {
	Conversations []*ConversationDataDto
	CompanyID     int64
}

type BulkMessageParamsDto struct {
	Messages  []*MessageDataDto
	CompanyID int64
}

type ConversationDataDto struct {
	JID               string
	Name              string
	UnreadCount       int32
	IsGroup           bool
	ProfilePictureURL string
}

type MessageDataDto struct {
	ConversationJID string
	RemoteJID       string
	FromMe          bool
	MessageText     string
	MediaURL        string
	MediaCaption    string
	Timestamp       uint64
	ReceivedAt      time.Time
	EditedAt        time.Time
	IsForwarded     bool
	IsDeleted       bool
}
