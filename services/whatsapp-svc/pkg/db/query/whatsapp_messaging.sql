-- name: CreateConversation :one
INSERT INTO 
  whatsapp_messaging.whatsapp_conversations (
	  id, 
    company_id, 
    user_id, 
    jid, 
    name,
    unread_count,
    is_group,
    profile_picture_url
	) 
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
  )
RETURNING id;

-- name: GetConversationByJID :one
SELECT
  id,
  company_id,
  user_id,
  jid,
  name,
  unread_count,
  is_group,
  profile_picture_url,
  last_message_timestamp
FROM
  whatsapp_messaging.whatsapp_conversations
WHERE
  company_id = $1 AND 
  user_id = $2 AND
  jid = $3
LIMIT 1;

-- name: GetConversationsByUser :many
SELECT
  id,
  company_id,
  user_id,
  jid,
  name,
  unread_count,
  is_group,
  profile_picture_url,
  last_message_timestamp
FROM
  whatsapp_messaging.whatsapp_conversations
WHERE
  company_id = $1 AND 
  user_id = $2
ORDER BY
  last_message_timestamp DESC NULLS LAST
LIMIT 
  $3
OFFSET 
  $4;

-- name: CreateMessage :one
INSERT INTO 
  whatsapp_messaging.whatsapp_messages (
	  id, 
    company_id, 
    conversation_id, 
    message_id, 
    remote_jid,
    from_me,
    message_type,
    message_text,
    media_url,
    media_caption,
    status,
    timestamp,
    received_at,
    edited_at,
    is_forwarded,
    is_deleted
	) 
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
  )
RETURNING id;

-- name: GetMessagesByConversation :many
SELECT
  id, 
  company_id, 
  conversation_id, 
  message_id, 
  remote_jid,
  from_me,
  message_type,
  message_text,
  media_url,
  media_caption,
  status,
  timestamp,
  received_at,
  edited_at,
  is_forwarded,
  is_deleted
FROM
  whatsapp_messaging.whatsapp_messages
WHERE
  conversation_id = $1
ORDER BY
  received_at DESC
LIMIT 
  $2
OFFSET 
  $3;
