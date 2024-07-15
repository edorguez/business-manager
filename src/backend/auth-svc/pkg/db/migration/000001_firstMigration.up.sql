CREATE SCHEMA "auth";

CREATE TABLE "auth"."role" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "description" varchar(100),
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "modified_at" timestamptz NOT NULL DEFAULT (NOW())
);

CREATE TABLE "auth"."user" (
  "id" bigserial PRIMARY KEY,
  "company_id" bigint NOT NULL,
  "role_id" bigint NOT NULL,
  "email" varchar(100) NOT NULL,
  "password_hash" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "modified_at" timestamptz NOT NULL DEFAULT (NOW())
);

ALTER TABLE "auth"."user" ADD FOREIGN KEY ("role_id") REFERENCES "auth"."role" ("id");

INSERT INTO "auth"."role" (id, name, description, created_at, modified_at) VALUES (1, 'SuperAdmin', 'Super Admin user with all permission', NOW(), NOW());
INSERT INTO "auth"."role" (id, name, description, created_at, modified_at) VALUES (2, 'Admin', 'Admin user', NOW(), NOW());
INSERT INTO "auth"."role" (id, name, description, created_at, modified_at) VALUES (3, 'Regular', 'Regular user', NOW(), NOW());
