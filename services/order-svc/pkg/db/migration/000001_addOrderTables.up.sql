CREATE SCHEMA "order";

CREATE TABLE "order"."order" (
  "id" bigserial PRIMARY KEY,
  "company_id" bigint NOT NULL,
  "customer_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "modified_at" timestamptz NOT NULL DEFAULT (NOW())
);

CREATE TABLE "order"."order_product" (
  "id" bigserial PRIMARY KEY,
  "order_id" bigint NOT NULL,
  "product_id" varchar(24) NOT NULL,
  "quantity" integer NOT NULL,
  "price" bigint NOT NULL,
  "name" varchar(50) NOT NULL,
  "image_url" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (NOW()),
  "modified_at" timestamptz NOT NULL DEFAULT (NOW())
);
