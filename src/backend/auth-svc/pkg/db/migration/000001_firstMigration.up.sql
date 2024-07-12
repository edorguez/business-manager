CREATE SCHEMA "auth";

CREATE TABLE "auth"."user" (
  "id" bigserial PRIMARY KEY,
  "email" varchar(100) NOT NULL,
  "password" varchar(20) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "modified_at" timestamptz NOT NULL DEFAULT (NOW())
);
