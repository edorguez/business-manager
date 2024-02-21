CREATE SCHEMA "client";

CREATE TABLE "client"."client" (
  "id" bigserial PRIMARY KEY,
  "company_id" bigint NOT NULL,
  "first_name" varchar(20) NOT NULL,
  "last_name" varchar(20),
  "email" varchar(100),
  "phone" varchar(11),
  "identification_number" varchar(20) NOT NULL,
  "identification_type" varchar(1) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);