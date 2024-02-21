-- name: CreateClient :one
INSERT INTO 
  client.client (
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

-- name: GetClient :one
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
  client.client
WHERE 
  id = $1 
LIMIT 1;

-- name: GetClients :many
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
  client.client
ORDER BY 
  id
LIMIT 
  $1
OFFSET 
  $2;

-- name: UpdateClient :one
UPDATE 
  client.client
SET 
  first_name = $2,
  modified_at = NOW()
WHERE 
  id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM 
  client.client
WHERE 
  id = $1;