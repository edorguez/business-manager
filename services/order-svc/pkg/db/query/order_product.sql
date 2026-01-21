-- name: CreateOrderProduct :one
INSERT INTO "order"."order_product" (order_id, product_id, quantity, price, name, image_url)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetOrderProductsByOrderId :many
SELECT * FROM "order"."order_product" WHERE order_id = $1 ORDER BY id;
