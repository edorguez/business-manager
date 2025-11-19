-- name: CreateBusinessPhone :one
INSERT INTO
  whatsapp.business_phone (
    company_id,
    phone
  )
VALUES (
  $1, $2
)
RETURNING *;

-- name: GetBusinessPhone :one
SELECT
  id,
  company_id,
  phone,
  created_at,
  modified_at
FROM
  whatsapp.business_phone
WHERE
  LOWER(phone) = LOWER($1)
LIMIT 1;

-- name: GetBusinessPhoneByCompanyId :one
SELECT
  id,
  company_id,
  phone,
  created_at,
  modified_at
FROM
  whatsapp.business_phone
WHERE
  company_id = $1
LIMIT 1;

-- name: UpdateBusinessPhone :one
UPDATE
  whatsapp.business_phone
SET
  phone = $1,
  modified_at = NOW()
WHERE
  company_id = $2
RETURNING *;
