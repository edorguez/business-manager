CREATE SCHEMA IF NOT EXISTS "whatsapp_messaging";

CREATE TABLE IF NOT EXISTS "whatsapp_messaging"."whatsapp_conversations" (
    "id" bigserial PRIMARY KEY,
    "company_id" bigint NOT NULL,
    "jid" varchar(100) NOT NULL,
    "name" VARCHAR(255),
    "unread_count" INTEGER DEFAULT 0,
    "is_group" BOOLEAN DEFAULT FALSE,
    "profile_picture_url" TEXT,
    "last_message_timestamp" timestamptz,
    "created_at" timestamptz DEFAULT (NOW()),
    "modified_at" timestamptz DEFAULT (NOW()),
    UNIQUE("jid")
);

CREATE TABLE IF NOT EXISTS "whatsapp_messaging"."whatsapp_messages" (
    "id" bigserial PRIMARY KEY,
    "company_id" bigint NOT NULL,
    "conversation_jid" varchar(100) NOT NULL REFERENCES "whatsapp_messaging"."whatsapp_conversations"("jid"),
    "remote_jid" varchar(100) NOT NULL,
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
    CONSTRAINT fk_conversation FOREIGN KEY ("conversation_jid") REFERENCES "whatsapp_messaging"."whatsapp_conversations"("jid")
);

CREATE INDEX IF NOT EXISTS idx_whatsapp_conversations_jid ON "whatsapp_messaging"."whatsapp_conversations"("jid"); 
CREATE INDEX IF NOT EXISTS idx_whatsapp_messages_conversation_timestamp ON "whatsapp_messaging"."whatsapp_messages"("conversation_jid", "timestamp");
CREATE INDEX IF NOT EXISTS idx_whatsapp_messages_tenant_conversation ON "whatsapp_messaging"."whatsapp_messages"("company_id", "conversation_jid");
