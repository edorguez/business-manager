ALTER TABLE "order"."order_product" 
ADD COLUMN "quantity" integer NOT NULL DEFAULT 1,
ADD COLUMN "price" bigint NOT NULL DEFAULT 0,
ADD COLUMN "name" varchar NOT NULL DEFAULT '';