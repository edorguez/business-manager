CREATE SCHEMA IF NOT EXISTS "whatsapp_messaging";

CREATE TABLE IF NOT EXISTS "whatsapp_messaging"."whatsapp_users" (
    "id" bigserial PRIMARY KEY,
    "company_id" bigint NOT NULL,
    "phone" varchar(50) NOT NULL,
    "whatsapp_jid" varchar(255) NOT NULL,
    "display_name" varchar(255),
    "profile_picture_url" text,
    "is_connected" boolean DEFAULT FALSE,
    "last_connected_at" timestamptz NOT NULL DEFAULT (NOW()),
    "created_at" timestamptz DEFAULT (NOW()),
    "modified_at" timestamptz DEFAULT (NOW()),
    UNIQUE("company_id", "phone"),
    UNIQUE("whatsapp_jid")
);

CREATE TABLE IF NOT EXISTS "whatsapp_messaging"."whatsapp_conversations" (
    "id" bigserial PRIMARY KEY,
    "company_id" bigint NOT NULL,
    "user_id" bigint NOT NULL REFERENCES "whatsapp_messaging"."whatsapp_users"("id") ON DELETE CASCADE,
    "jid" varchar(255) NOT NULL,
    "name" VARCHAR(255),
    "unread_count" INTEGER DEFAULT 0,
    "is_group" BOOLEAN DEFAULT FALSE,
    "profile_picture_url" TEXT,
    "last_message_timestamp" timestamptz,
    "created_at" timestamptz DEFAULT (NOW()),
    "modified_at" timestamptz DEFAULT (NOW()),
    UNIQUE("company_id", "user_id", "jid")
);

CREATE TABLE IF NOT EXISTS "whatsapp_messaging"."whatsapp_messages" (
    "id" bigserial PRIMARY KEY,
    "company_id" bigint NOT NULL,
    "conversation_id" bigint NOT NULL REFERENCES "whatsapp_messaging"."whatsapp_conversations"("id") ON DELETE CASCADE,
    "message_id" varchar(255) NOT NULL,
    "remote_jid" varchar(255) NOT NULL,
    "from_me" boolean DEFAULT FALSE,
    "message_type" varchar(50) NOT NULL,
    "message_text" TEXT,
    "media_url" TEXT,
    "media_caption" TEXT,
    "status" varchar(50) DEFAULT 'sent',
    "timestamp" TIMESTAMP WITH TIME ZONE NOT NULL,
    "received_at" timestamptz,
    "edited_at" timestamptz,
    "is_forwarded" boolean DEFAULT FALSE,
    "is_deleted" boolean DEFAULT FALSE,
    "created_at" timestamptz DEFAULT (NOW()),
    "modified_at" timestamptz DEFAULT (NOW()),
    UNIQUE("company_id", "conversation_id", "message_id")
);

CREATE INDEX idx_whatsapp_conversations_tenant_user ON "whatsapp_messaging"."whatsapp_conversations"("company_id", "user_id");
CREATE INDEX idx_whatsapp_conversations_jid ON "whatsapp_messaging"."whatsapp_conversations"("jid");
CREATE INDEX idx_whatsapp_messages_conversation_timestamp ON "whatsapp_messaging"."whatsapp_messages"("conversation_id", "timestamp");
CREATE INDEX idx_whatsapp_messages_tenant_conversation ON "whatsapp_messaging"."whatsapp_messages"("company_id", "conversation_id");
CREATE INDEX idx_whatsapp_messages_message_id ON "whatsapp_messaging"."whatsapp_messages"("message_id");
