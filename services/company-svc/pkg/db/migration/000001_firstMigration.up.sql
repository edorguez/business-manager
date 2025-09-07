CREATE SCHEMA "company";

CREATE TABLE "company"."company" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "name_format_url" varchar NOT NULL,
  "image_url" varchar,
  "plan_id" bigint NOT NULL,
  "last_payment_date" timestamptz NOT NULL DEFAULT (NOW()),
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "modified_at" timestamptz NOT NULL DEFAULT (NOW())
);

CREATE TABLE "company"."payment" (
  "id" bigserial PRIMARY KEY,
  "company_id" bigint NOT NULL,
  "name" varchar(50) NOT NULL,
  "bank" varchar(50),
  "account_number" varchar(20),
  "account_type" varchar(20),
  "identification_number" varchar(9),
  "identification_type" varchar(1),
  "phone" varchar(11),
  "email" varchar(100),
  "payment_type_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "modified_at" timestamptz NOT NULL DEFAULT (NOW())
);

CREATE TABLE "company"."payment_type" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "modified_at" timestamptz NOT NULL DEFAULT (NOW())
);

CREATE TABLE "company"."plan" (
    "id" bigserial PRIMARY KEY,
    "name" varchar(50) NOT NULL,
    "description" varchar(100),
    "created_at" timestamptz NOT NULL DEFAULT (NOW()),
    "modified_at" timestamptz NOT NULL DEFAULT (NOW())
);

ALTER TABLE "company"."payment" ADD FOREIGN KEY ("payment_type_id") REFERENCES "company"."payment_type" ("id");
ALTER TABLE "company"."company" ADD FOREIGN KEY ("plan_id") REFERENCES "company"."plan" ("id");

INSERT INTO "company"."payment_type" (id, name, created_at, modified_at) VALUES (1, 'Binance', NOW(), NOW());
INSERT INTO "company"."payment_type" (id, name, created_at, modified_at) VALUES (2, 'Colombia', NOW(), NOW());
INSERT INTO "company"."payment_type" (id, name, created_at, modified_at) VALUES (3, 'Pago Móvil', NOW(), NOW());
INSERT INTO "company"."payment_type" (id, name, created_at, modified_at) VALUES (4, 'Otro', NOW(), NOW());
INSERT INTO "company"."payment_type" (id, name, created_at, modified_at) VALUES (5, 'Panamá', NOW(), NOW());
INSERT INTO "company"."payment_type" (id, name, created_at, modified_at) VALUES (6, 'Paypal', NOW(), NOW());
INSERT INTO "company"."payment_type" (id, name, created_at, modified_at) VALUES (7, 'Transferencia', NOW(), NOW());
INSERT INTO "company"."payment_type" (id, name, created_at, modified_at) VALUES (8, 'USA', NOW(), NOW());
INSERT INTO "company"."payment_type" (id, name, created_at, modified_at) VALUES (9, 'Zelle', NOW(), NOW());
INSERT INTO "company"."payment_type" (id, name, created_at, modified_at) VALUES (10, 'Zinli', NOW(), NOW());

INSERT INTO "company"."plan" (id, name, description, created_at, modified_at) VALUES (1, 'Basic', 'Plan with basic features', NOW(), NOW());
INSERT INTO "company"."plan" (id, name, description, created_at, modified_at) VALUES (2, 'Pro', 'Plan with pro features', NOW(), NOW());