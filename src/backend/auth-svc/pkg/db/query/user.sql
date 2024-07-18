-- name: CreateUser :one
INSERT INTO 
  auth.user (
    company_id,
    role_id,
    email,
    password_hash
  ) 
VALUES (
  $1, $2, $3, $4
) 
RETURNING *;

-- name: GetUser :one
SELECT 
  id,
  company_id,
  role_id,
  email,
  password_hash,
  created_at,
  modified_at
FROM 
  auth.user
WHERE 
  id = $1 
LIMIT 1;

-- name: GetUserByEmail :one
SELECT 
  id,
  company_id,
  role_id,
  email,
  password_hash,
  created_at,
  modified_at
FROM 
  auth.user
WHERE 
  email = $1 
LIMIT 1;


-- name: GetUsers :many
SELECT 
  id,
  company_id,
  role_id,
  email,
  password_hash,
  created_at,
  modified_at
FROM 
  auth.user
WHERE
  (company_id = $1) OR $1 = 0
ORDER BY 
  id
LIMIT 
  $2
OFFSET 
  $3;

-- name: UpdateUser :one
UPDATE 
  auth.user
SET 
  role_id = $2,
  email = $3,
  password_hash = $4,
  modified_at = NOW()
WHERE 
  id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM 
  auth.user
WHERE 
  id = $1;
