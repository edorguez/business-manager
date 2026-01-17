-- name: CreateOrder :one
INSERT INTO "order"."order" (company_id, customer_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM "order"."order" WHERE id = $1;

-- name: GetOrders :many
SELECT * FROM "order"."order"
WHERE company_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetOrdersCount :one
SELECT COUNT(*) FROM "order"."order" WHERE company_id = $1;