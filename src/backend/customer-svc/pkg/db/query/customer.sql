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
  (@company_id = 0 OR company_id = @company_id) AND
  (@first_name::text = '' OR first_name LIKE CONCAT('%', @first_name::text, '%')) AND
  (@last_name::text = '' OR last_name LIKE CONCAT('%', @last_name::text, '%')) AND
  (@identification_number::text = '' OR identification_number LIKE CONCAT('%', @identification_number::text, '%'))
ORDER BY 
  id
LIMIT 
  $1
OFFSET 
  $2;

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

-- name: GetCustomersByMonths :many
SELECT 
    DATE_TRUNC('month', created_at)::timestamp AS month_interval,
    COUNT(*) AS record_count
FROM 
    customer.customer
WHERE
  @company_id = 0 OR company_id = @company_id
GROUP BY 
    month_interval
ORDER BY 
    month_interval;