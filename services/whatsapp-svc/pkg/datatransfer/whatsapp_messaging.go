package datatransfer

import (
	"time"
)

type CreateConversationRequestDto struct {
	ID                int64   `json:"id"`
	CompanyID         int64   `json:"company_id"`
	UserID            int64   `json:"user_id"`
	Jid               string  `json:"jid"`
	Name              *string `json:"name"`
	UnreadCount       *int32  `json:"unread_count"`
	IsGroup           *bool   `json:"is_group"`
	ProfilePictureUrl *string `json:"profile_picture_url"`
}

type GetConversationByJIDRequestDto struct {
	CompanyID int64  `json:"company_id"`
	UserID    int64  `json:"user_id"`
	Jid       string `json:"jid"`
}

type GetConversationByJIDResponseDto struct {
	ID                   int64      `json:"id"`
	CompanyID            int64      `json:"company_id"`
	UserID               int64      `json:"user_id"`
	Jid                  string     `json:"jid"`
	Name                 *string    `json:"name"`
	UnreadCount          *int32     `json:"unread_count"`
	IsGroup              *bool      `json:"is_group"`
	ProfilePictureUrl    *string    `json:"profile_picture_url"`
	LastMessageTimestamp *time.Time `json:"last_message_timestamp"`
}

type GetConversationsByUserRequestDto struct {
	CompanyID int64 `json:"company_id"`
	UserID    int64 `json:"user_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

type GetConversationsByUserResponseDto struct {
	ID                   int64      `json:"id"`
	CompanyID            int64      `json:"company_id"`
	UserID               int64      `json:"user_id"`
	Jid                  string     `json:"jid"`
	Name                 *string    `json:"name"`
	UnreadCount          *int32     `json:"unread_count"`
	IsGroup              *bool      `json:"is_group"`
	ProfilePictureUrl    *string    `json:"profile_picture_url"`
	LastMessageTimestamp *time.Time `json:"last_message_timestamp"`
}

type CreateMessageRequestDto struct {
	ID             int64      `json:"id"`
	CompanyID      int64      `json:"company_id"`
	ConversationID int64      `json:"conversation_id"`
	MessageID      string     `json:"message_id"`
	RemoteJid      string     `json:"remote_jid"`
	FromMe         *bool      `json:"from_me"`
	MessageType    string     `json:"message_type"`
	MessageText    *string    `json:"message_text"`
	MediaUrl       *string    `json:"media_url"`
	MediaCaption   *string    `json:"media_caption"`
	Status         *string    `json:"status"`
	Timestamp      time.Time  `json:"timestamp"`
	ReceivedAt     *time.Time `json:"received_at"`
	EditedAt       *time.Time `json:"edited_at"`
	IsForwarded    *bool      `json:"is_forwarded"`
	IsDeleted      *bool      `json:"is_deleted"`
}

type GetMessagesByConversationRequestDto struct {
	ConversationID int64 `json:"conversation_id"`
	Limit          int32 `json:"limit"`
	Offset         int32 `json:"offset"`
}

type GetMessagesByConversationResponseDto struct {
	ID             int64      `json:"id"`
	CompanyID      int64      `json:"company_id"`
	ConversationID int64      `json:"conversation_id"`
	MessageID      string     `json:"message_id"`
	RemoteJid      string     `json:"remote_jid"`
	FromMe         *bool      `json:"from_me"`
	MessageType    string     `json:"message_type"`
	MessageText    *string    `json:"message_text"`
	MediaUrl       *string    `json:"media_url"`
	MediaCaption   *string    `json:"media_caption"`
	Status         *string    `json:"status"`
	Timestamp      time.Time  `json:"timestamp"`
	ReceivedAt     *time.Time `json:"received_at"`
	EditedAt       *time.Time `json:"edited_at"`
	IsForwarded    *bool      `json:"is_forwarded"`
	IsDeleted      *bool      `json:"is_deleted"`
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
	ID                int64
	UserID            int64
	JID               string
	Name              string
	UnreadCount       int32
	IsGroup           bool
	ProfilePictureURL string
}

type MessageDataDto struct {
	ID             int64
	ConversationID int64
	MessageID      string
	RemoteJID      string
	FromMe         bool
	MessageType    string
	MessageText    string
	MediaURL       string
	MediaCaption   string
	Status         string
	Timestamp      time.Time
	ReceivedAt     time.Time
	EditedAt       time.Time
	IsForwarded    bool
	IsDeleted      bool
}
