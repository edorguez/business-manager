CREATE SCHEMA "whatsapp";

CREATE TABLE "whatsapp"."business_phone" (
  "id" bigserial PRIMARY KEY,
  "company_id" bigint NOT NULL,
  "phone" varchar(11),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);
