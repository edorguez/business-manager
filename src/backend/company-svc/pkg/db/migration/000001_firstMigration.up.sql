CREATE SCHEMA "company";

CREATE TABLE "company"."company" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "image_url" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "company"."payment" (
  "id" bigserial PRIMARY KEY,
  "company_id" bigint NOT NULL,
  "bank" varchar(50),
  "account_number" varchar(20),
  "account_type" varchar(20),
  "identification_number" varchar(20),
  "identification_type" varchar(1),
  "phone" varchar(11),
  "email" varchar(100),
  "payment_type_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "company"."payment_type" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "modified_at" timestamptz NOT NULL DEFAULT (now())
);
