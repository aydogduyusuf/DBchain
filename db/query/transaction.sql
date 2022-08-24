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
WHERE 
    id = $1
LIMIT 1;


-- name: ListTransactions :many
SELECT * FROM transactions
WHERE 
    from_address = $1 OR
    to_address = $2
ORDER BY id;

-- name: DeleteTransaction :exec
UPDATE transactions
SET is_active = false AND delete_time = current_timestamp
WHERE id = $1;