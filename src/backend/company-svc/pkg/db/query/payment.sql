-- name: CreatePayment :one
INSERT INTO 
  company.payment (
    company_id,
    name,
    bank,
    account_number,
    account_type,
    identification_number,
    identification_type,
    phone,
    email,
    payment_type_id
  ) 
VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10 
) 
RETURNING *;

-- name: GetPayment :one
SELECT 
  P.id,
  P.company_id,
  P.name,
  P.bank,
  P.account_number,
  P.account_type,
  P.identification_number,
  P.identification_type,
  P.phone,
  P.email,
  P.payment_type_id,
  P.created_at,
  P.modified_at,
  sqlc.embed(PT)
FROM 
  company.payment AS P
INNER JOIN 
  company.payment_type AS PT ON P.payment_type_id = PT.id
WHERE 
  P.id = $1 
LIMIT 1;

-- name: GetPayments :many
SELECT 
  P.id,
  P.company_id,
  P.name,
  P.bank,
  P.account_number,
  P.account_type,
  P.identification_number,
  P.identification_type,
  P.phone,
  P.email,
  P.payment_type_id,
  P.created_at,
  P.modified_at,
  sqlc.embed(PT)
FROM 
  company.payment AS P
INNER JOIN 
  company.payment_type AS PT ON P.payment_type_id = PT.id
WHERE
  (P.company_id = $1) OR $1 = 0
ORDER BY 
  P.id
LIMIT 
  $2
OFFSET 
  $3;

-- name: UpdatePayment :one
UPDATE 
  company.payment
SET 
  name = $2,
  bank = $3,
  account_number = $4,
  account_type = $5,
  identification_number = $6,
  identification_type = $7,
  phone = $8,
  email = $9,
  payment_type_id = $10,  
  modified_at = NOW()
WHERE 
  id = $1
RETURNING *;

-- name: DeletePayment :exec
DELETE FROM 
  company.payment
WHERE 
  id = $1;
