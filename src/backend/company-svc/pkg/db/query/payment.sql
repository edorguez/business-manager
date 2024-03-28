-- name: CreatePayment :one
INSERT INTO 
  company.payment (
    name,
    image_url
  ) 
VALUES (
  $1, $2 
) 
RETURNING *;

-- name: GetPayment :one
SELECT 
  id,
  name,
  image_url,
  created_at,
  modified_at
FROM 
  company.payment
WHERE 
  id = $1 
LIMIT 1;

-- name: GetPayments :many
SELECT 
  id,
  name,
  image_url,
  created_at,
  modified_at
FROM 
  company.payment
ORDER BY 
  id
LIMIT 
  $1
OFFSET 
  $2;

-- name: UpdatePayment :one
UPDATE 
  company.payment
SET 
  name = $2,
  image_url = $3,
  modified_at = NOW()
WHERE 
  id = $1
RETURNING *;

-- name: DeletePayment :exec
DELETE FROM 
  company.payment
WHERE 
  id = $1;
