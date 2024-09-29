-- name: GetRole :one
SELECT 
  id,
  name,
  description
FROM 
  auth.role
WHERE 
  id = $1 
LIMIT 1;

-- name: GetRoles :many
SELECT 
  id,
  name,
  description
FROM 
  auth.role
ORDER BY 
  id;