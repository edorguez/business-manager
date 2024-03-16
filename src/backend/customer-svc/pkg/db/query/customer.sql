-- name: CreateCustomer :one
INSERT INTO 
  customer.customer (
    company_id,
    first_name,
    last_name,
    email,
    phone,
    identification_number,
    identification_type
  ) 
VALUES (
  $1, $2, $3, $4, $5, $6, $7
) 
RETURNING *;

-- name: GetCustomer :one
SELECT 
  id,
  company_id,
  first_name,
  last_name,
  email,
  phone,
  identification_number,
  identification_type,
  created_at,
  modified_at
FROM 
  customer.customer
WHERE 
  id = $1 
LIMIT 1;

-- name: GetCustomers :many
SELECT 
  id,
  company_id,
  first_name,
  last_name,
  email,
  phone,
  identification_number,
  identification_type,
  created_at,
  modified_at
FROM 
  customer.customer
WHERE
  (company_id = $1) OR $1 = 0
ORDER BY 
  id
LIMIT 
  $2
OFFSET 
  $3;

-- name: UpdateCustomer :one
UPDATE 
  customer.customer
SET 
  first_name = $2,
  last_name = $3,
  email = $4,
  phone = $5,
  identification_number = $6,
  identification_type = $7,
  modified_at = NOW()
WHERE 
  id = $1
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM 
  customer.customer
WHERE 
  id = $1;
