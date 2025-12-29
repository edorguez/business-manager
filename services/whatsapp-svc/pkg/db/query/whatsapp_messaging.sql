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

-- name: BulkUpsertConversations :exec
INSERT INTO whatsapp_messaging.whatsapp_conversations (
    id, company_id, user_id, jid, name, unread_count, is_group, profile_picture_url
) 
SELECT 
    unnest(@ids::bigint[]) as id,
    unnest(@company_ids::bigint[]) as company_id,
    unnest(@user_ids::bigint[]) as user_id,
    unnest(@jids::text[]) as jid,
    unnest(@names::text[]) as name,
    unnest(@unread_counts::int[]) as unread_count,
    unnest(@is_groups::boolean[]) as is_group,
    unnest(@profile_picture_urls::text[]) as profile_picture_url
ON CONFLICT (company_id, user_id, jid) 
DO UPDATE SET
    name = EXCLUDED.name,
    unread_count = EXCLUDED.unread_count,
    is_group = EXCLUDED.is_group,
    profile_picture_url = EXCLUDED.profile_picture_url,
    last_message_timestamp = CASE 
        WHEN EXCLUDED.last_message_timestamp > whatsapp_messaging.whatsapp_conversations.last_message_timestamp 
        OR whatsapp_messaging.whatsapp_conversations.last_message_timestamp IS NULL
        THEN EXCLUDED.last_message_timestamp
        ELSE whatsapp_messaging.whatsapp_conversations.last_message_timestamp
    END
WHERE excluded.id IS NOT NULL;

-- name: BulkUpsertMessages :exec
INSERT INTO whatsapp_messaging.whatsapp_messages (
    id, company_id, conversation_id, message_id, remote_jid, from_me,
    message_type, message_text, media_url, media_caption, status,
    timestamp, received_at, edited_at, is_forwarded, is_deleted
) 
SELECT 
    unnest(@ids::bigint[]) as id,
    unnest(@company_ids::bigint[]) as company_id,
    unnest(@conversation_ids::bigint[]) as conversation_id,
    unnest(@message_ids::text[]) as message_id,
    unnest(@remote_jids::text[]) as remote_jid,
    unnest(@from_mes::boolean[]) as from_me,
    unnest(@message_types::text[]) as message_type,
    unnest(@message_texts::text[]) as message_text,
    unnest(@media_urls::text[]) as media_url,
    unnest(@media_captions::text[]) as media_caption,
    unnest(@statuses::text[]) as status,
    unnest(@timestamps::timestamptz[]) as timestamp,
    unnest(@received_ats::timestamptz[]) as received_at,
    unnest(@edited_ats::timestamptz[]) as edited_at,
    unnest(@is_forwardeds::boolean[]) as is_forwarded,
    unnest(@is_deleteds::boolean[]) as is_deleted
ON CONFLICT (company_id, conversation_id, message_id) 
DO UPDATE SET
    message_text = EXCLUDED.message_text,
    media_url = EXCLUDED.media_url,
    media_caption = EXCLUDED.media_caption,
    status = EXCLUDED.status,
    edited_at = EXCLUDED.edited_at,
    is_forwarded = EXCLUDED.is_forwarded,
    is_deleted = EXCLUDED.is_deleted
WHERE excluded.id IS NOT NULL;

-- name: GetConversationByID :one
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
  id = $1
LIMIT 1;
