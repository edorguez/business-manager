-- name: CreateClient :one
INSERT INTO client (
  company_id,
  first_name,
  last_name,
  email,
  phone,
  identification_number,
  identification_type
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;