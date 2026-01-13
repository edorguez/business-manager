-- name: CreateConversation :one
INSERT INTO 
  whatsapp_messaging.whatsapp_conversations (
    company_id, 
    jid, 
    name,
    unread_count,
    is_group,
    profile_picture_url
	) 
VALUES (
  $1, $2, $3, $4, $5, $6
  )
RETURNING id;

-- name: GetConversationByJID :one
SELECT
  id,
  company_id,
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
  jid = $2
LIMIT 1;

-- name: CreateMessage :one
INSERT INTO 
  whatsapp_messaging.whatsapp_messages (
    company_id, 
    conversation_jid, 
    remote_jid,
    from_me,
    message_text,
    media_url,
    media_caption,
    timestamp,
    received_at,
    edited_at,
    is_forwarded,
    is_deleted
	) 
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
  )
RETURNING id;

-- name: GetMessagesByConversationJID :many
SELECT
  id, 
  company_id, 
  conversation_jid, 
  remote_jid,
  from_me,
  message_text,
  media_url,
  media_caption,
  timestamp,
  received_at,
  edited_at,
  is_forwarded,
  is_deleted
FROM
  whatsapp_messaging.whatsapp_messages
WHERE
  conversation_jid = $1
ORDER BY
  received_at DESC
LIMIT 
  $2
OFFSET 
  $3;

-- name: BulkUpsertConversations :exec
INSERT INTO whatsapp_messaging.whatsapp_conversations (
    company_id, jid, name, unread_count, is_group, profile_picture_url
) 
SELECT 
    unnest(@company_ids::bigint[]) as company_id,
    unnest(@jids::text[]) as jid,
    unnest(@names::text[]) as name,
    unnest(@unread_counts::int[]) as unread_count,
    unnest(@is_groups::boolean[]) as is_group,
    unnest(@profile_picture_urls::text[]) as profile_picture_url
ON CONFLICT (company_id, jid) 
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
    END,
    modified_at = NOW();

-- name: BulkUpsertMessages :exec
WITH input_data AS (
    SELECT DISTINCT ON (company_id, conversation_jid, timestamp, from_me, remote_jid)
        company_id, conversation_jid, remote_jid, from_me,
        message_text, media_url, media_caption,
        timestamp, received_at, edited_at, is_forwarded, is_deleted
    FROM (
        SELECT 
            unnest(@company_ids::bigint[]) as company_id,
            unnest(@conversation_jids::text[]) as conversation_jid,
            unnest(@remote_jids::text[]) as remote_jid,
            unnest(@from_mes::boolean[]) as from_me,
            unnest(@message_texts::text[]) as message_text,
            unnest(@media_urls::text[]) as media_url,
            unnest(@media_captions::text[]) as media_caption,
            unnest(@timestamps::bigint[]) as timestamp,
            unnest(@received_ats::timestamptz[]) as received_at,
            unnest(@edited_ats::timestamptz[]) as edited_at,
            unnest(@is_forwardeds::boolean[]) as is_forwarded,
            unnest(@is_deleteds::boolean[]) as is_deleted
    ) AS raw_input
    ORDER BY company_id, conversation_jid, timestamp, from_me, remote_jid, timestamp DESC
)
INSERT INTO whatsapp_messaging.whatsapp_messages (
    company_id, conversation_jid, remote_jid, from_me,
    message_text, media_url, media_caption,
    timestamp, received_at, edited_at, is_forwarded, is_deleted
) 
SELECT * FROM input_data
ON CONFLICT (company_id, conversation_jid, timestamp, from_me, remote_jid)
DO UPDATE SET
    message_text = EXCLUDED.message_text,
    media_url = EXCLUDED.media_url,
    media_caption = EXCLUDED.media_caption,
    edited_at = EXCLUDED.edited_at,
    is_forwarded = EXCLUDED.is_forwarded,
    is_deleted = EXCLUDED.is_deleted,
    modified_at = NOW();

-- name: GetConversationByID :one
SELECT
  id,
  company_id,
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
