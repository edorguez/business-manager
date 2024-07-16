-- name: GetRole :one
SELECT 
  id,
  name,
  description,
  created_at,
  modified_at
FROM 
  auth.role
WHERE 
  id = $1 
LIMIT 1;

-- name: GetRoles :many
SELECT 
  id,
  name,
  description,
  created_at,
  modified_at
FROM 
  auth.role
ORDER BY 
  id
LIMIT 
  $1
OFFSET 
  $2;

