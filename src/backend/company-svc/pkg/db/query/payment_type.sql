-- name: GetPaymentType :one
SELECT 
  id,
  name,
  created_at,
  modified_at
FROM 
  company.payment_type
WHERE 
  id = $1 
LIMIT 1;

-- name: GetPaymentTypes :many
SELECT 
  id,
  name,
  created_at,
  modified_at
FROM 
  company.payment_type
LIMIT 
  $1
OFFSET 
  $2;