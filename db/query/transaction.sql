-- name: CreateTransaction :one
INSERT INTO transactions (
    transaction_type,
    from_address,
    to_address,
    transfer_data,
    hash_value
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;


-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1
LIMIT 1;

-- name: GetTransactionByAddress :one
SELECT * FROM transactions
WHERE transfer_data = $1
LIMIT 1;


-- name: ListTransactionsByTypeFrom :many
SELECT * FROM transactions
WHERE 
    from_address = $1 AND transaction_type = $2
ORDER BY id;

-- name: ListTransactionsByTypeTo :many
SELECT * FROM transactions
WHERE 
    to_address = $1 AND transaction_type = $2
ORDER BY id;

-- name: ListTransactionsByToken :many
SELECT * FROM transactions
WHERE 
    transfer_data = $1
ORDER BY id;

-- name: ListDeploysByUser :many
SELECT * FROM transactions
WHERE 
    from_address = $1 AND transaction_type = $2
ORDER BY id;

-- name: ListTransfersByTimeFrom :many
SELECT * FROM transactions
WHERE 
    create_time >= $1 AND create_time <= $2 AND transaction_type = $3 AND from_address = $4
ORDER BY id;

-- name: ListTransfersByTimeTo :many
SELECT * FROM transactions
WHERE 
    create_time >= $1 AND create_time <= $2 AND transaction_type = $3 AND to_address = $4
ORDER BY id;

-- name: ListDeploysByTime :many
SELECT * FROM transactions
WHERE 
    create_time >= $1 AND create_time <= $2 AND transaction_type = $3 AND from_address = $4
ORDER BY id;


-- name: DeleteTransaction :exec
UPDATE transactions
SET is_active = false AND delete_time = current_timestamp
WHERE id = $1;