-- name: LockCompanyOrders :exec
SELECT pg_advisory_xact_lock($1);

-- name: GetMaxOrderNumberForCompany :one
SELECT COALESCE(MAX(order_number), 0)::INTEGER as max_order_number 
FROM "order"."order" 
WHERE company_id = $1;

-- name: CreateOrder :one
INSERT INTO "order"."order" (company_id, customer_id, order_number)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM "order"."order" WHERE id = $1;

-- name: GetOrders :many
SELECT * FROM "order"."order"
WHERE company_id = $1
ORDER BY created_at DESC, id DESC
LIMIT $2 OFFSET $3;

-- name: GetOrdersCount :one
SELECT COUNT(*) FROM "order"."order" WHERE company_id = $1;

-- name: GetOrdersByMonth :many
SELECT created_at
FROM "order"."order"
WHERE company_id = sqlc.arg(company_id)
  AND created_at >= make_date(sqlc.arg(year)::int, sqlc.arg(month)::int, 1)
  AND created_at < make_date(sqlc.arg(year)::int, sqlc.arg(month)::int, 1) + INTERVAL '1 month'
ORDER BY created_at;
