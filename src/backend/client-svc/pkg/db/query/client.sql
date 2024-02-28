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
WHERE
  (company_id = $1) OR $1 = 0
ORDER BY 
  id
LIMIT 
  $2
OFFSET 
  $3;

-- name: UpdateClient :one
UPDATE 
  client.client
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

-- name: DeleteClient :exec
DELETE FROM 
  client.client
WHERE 
  id = $1;