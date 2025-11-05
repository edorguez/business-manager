-- name: CreateCompany :one
INSERT INTO 
  company.company (
    name,
    name_format_url,
    is_free_trial,
    plan_id,
    last_payment_date
  ) 
VALUES (
  $1, $2, $3, $4, $5
) 
RETURNING *;

-- name: GetCompany :one
SELECT 
  id,
  name,
  name_format_url,
  image_url,
  plan_id,
  last_payment_date,
  created_at,
  modified_at
FROM 
  company.company
WHERE 
  id = $1 
LIMIT 1;

-- name: GetCompanyByName :one
SELECT 
  id,
  name,
  name_format_url,
  image_url,
  plan_id,
  last_payment_date,
  created_at,
  modified_at
FROM 
  company.company
WHERE 
  LOWER(name) = LOWER($1)
LIMIT 1;

-- name: GetCompanyByNameUrl :one
SELECT 
  id,
  name,
  name_format_url,
  image_url,
  plan_id,
  last_payment_date,
  created_at,
  modified_at
FROM 
  company.company
WHERE 
  LOWER(name_format_url) = LOWER($1)
LIMIT 1;

-- name: GetCompanies :many
SELECT 
  id,
  name,
  name_format_url,
  image_url,
  plan_id,
  last_payment_date,
  created_at,
  modified_at
FROM 
  company.company
ORDER BY 
  id
LIMIT 
  $1
OFFSET 
  $2;

-- name: UpdateCompany :one
UPDATE 
  company.company
SET 
  name = $2,
  name_format_url = $3,
  image_url = $4,
  modified_at = NOW()
WHERE 
  id = $1
RETURNING *;

-- name: DeleteCompany :exec
DELETE FROM 
  company.company
WHERE 
  id = $1;

-- name: UpdateCompanyImageUrl :one
UPDATE 
  company.company
SET 
  image_url = $2,
  modified_at = NOW()
WHERE 
  id = $1
RETURNING *;

