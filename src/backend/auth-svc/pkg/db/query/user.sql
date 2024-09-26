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
  U.id,
  U.company_id,
  U.role_id,
  U.email,
  sqlc.embed(R)
FROM 
  auth.user AS U
INNER JOIN
  auth.role AS R ON R.id = U.role_id
WHERE 
  U.id = $1 
LIMIT 1;

-- name: GetUserByEmail :one
SELECT 
  id,
  company_id,
  role_id,
  email,
  password_hash
FROM 
  auth.user
WHERE 
  email = $1 
LIMIT 1;

-- name: GetUsers :many
SELECT 
  U.id,
  U.company_id,
  U.role_id,
  U.email,
  sqlc.embed(R)
FROM 
  auth.user AS U
INNER JOIN
  auth.role AS R ON R.id = U.role_id
WHERE
  (U.company_id = $1) OR $1 = 0
ORDER BY 
  U.id
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

-- name: UpdateEmail :one
UPDATE 
  auth.user
SET 
  email = $2,
  modified_at = NOW()
WHERE 
  id = $1
RETURNING *;

-- name: UpdatePassword :one
UPDATE 
  auth.user
SET 
  password_hash = $2,
  modified_at = NOW()
WHERE 
  id = $1
RETURNING *;