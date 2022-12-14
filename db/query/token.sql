-- name: CreateToken :one
INSERT INTO tokens (
  u_id,
  token_name,
  symbol,
  supply,
  contract_address
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetToken :one
SELECT * FROM tokens
WHERE id = $1 LIMIT 1;

-- name: GetTokenByAddress :one
SELECT * FROM tokens
WHERE contract_address = $1 LIMIT 1;

-- name: GetTokenByUIDAndContract :one
SELECT * FROM tokens
WHERE u_id = $1 AND contract_address = $2 
LIMIT 1;

-- name: GetTokenForUpdate :one
SELECT * FROM tokens
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTokens :many
SELECT * FROM tokens
WHERE u_id = $1
ORDER BY id;

-- name: UpdateToken :one
UPDATE tokens 
SET contract_address = $2
WHERE id = $1
RETURNING *;

-- name: DeleteToken :exec
UPDATE tokens
SET is_active = false, delete_time = current_timestamp
WHERE id = $1;