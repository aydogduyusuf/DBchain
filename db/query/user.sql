-- name: CreateUser :one
INSERT INTO users (
  id,
  username,
  hashed_password,
  full_name,
  email,
  wallet_public_address,
  wallet_private_address
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: DeleteUser :exec
UPDATE users
SET status = false
WHERE id = $1;