ALTER TABLE "order"."order_product" 
DROP COLUMN IF EXISTS "quantity",
DROP COLUMN IF EXISTS "price",
DROP COLUMN IF EXISTS "name";