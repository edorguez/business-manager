ALTER TABLE "company"."payment_type" ADD COLUMN "image_path" VARCHAR(50);

UPDATE "company"."payment_type" SET "image_path" = 'binance.png' WHERE "id" = 1;
UPDATE "company"."payment_type" SET "image_path" = 'colombia.png' WHERE "id" = 2;
UPDATE "company"."payment_type" SET "image_path" = 'mobile.png' WHERE "id" = 3;
UPDATE "company"."payment_type" SET "image_path" = 'other.png' WHERE "id" = 4;
UPDATE "company"."payment_type" SET "image_path" = 'panama.png' WHERE "id" = 5;
UPDATE "company"."payment_type" SET "image_path" = 'paypal.png' WHERE "id" = 6;
UPDATE "company"."payment_type" SET "image_path" = 'transfer.png' WHERE "id" = 7;
UPDATE "company"."payment_type" SET "image_path" = 'usa.png' WHERE "id" = 8;
UPDATE "company"."payment_type" SET "image_path" = 'zelle.png' WHERE "id" = 9;
UPDATE "company"."payment_type" SET "image_path" = 'zinli.png' WHERE "id" = 10;