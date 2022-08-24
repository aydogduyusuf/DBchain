// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: token.sql

package db

import (
	"context"
)

const createToken = `-- name: CreateToken :one
INSERT INTO tokens (
  u_id,
  token_name,
  symbol,
  supply,
  contract_address
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, u_id, token_name, symbol, supply, contract_address, create_time, update_time, delete_time, is_active
`

type CreateTokenParams struct {
	UID             int64  `json:"u_id"`
	TokenName       string `json:"token_name"`
	Symbol          string `json:"symbol"`
	Supply          int64  `json:"supply"`
	ContractAddress string `json:"contract_address"`
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, createToken,
		arg.UID,
		arg.TokenName,
		arg.Symbol,
		arg.Supply,
		arg.ContractAddress,
	)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.UID,
		&i.TokenName,
		&i.Symbol,
		&i.Supply,
		&i.ContractAddress,
		&i.CreateTime,
		&i.UpdateTime,
		&i.DeleteTime,
		&i.IsActive,
	)
	return i, err
}

const deleteToken = `-- name: DeleteToken :exec
UPDATE tokens
SET is_active = false, delete_time = current_timestamp
WHERE id = $1
`

func (q *Queries) DeleteToken(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteToken, id)
	return err
}

const getToken = `-- name: GetToken :one
SELECT id, u_id, token_name, symbol, supply, contract_address, create_time, update_time, delete_time, is_active FROM tokens
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetToken(ctx context.Context, id int64) (Token, error) {
	row := q.db.QueryRowContext(ctx, getToken, id)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.UID,
		&i.TokenName,
		&i.Symbol,
		&i.Supply,
		&i.ContractAddress,
		&i.CreateTime,
		&i.UpdateTime,
		&i.DeleteTime,
		&i.IsActive,
	)
	return i, err
}

const getTokenForUpdate = `-- name: GetTokenForUpdate :one
SELECT id, u_id, token_name, symbol, supply, contract_address, create_time, update_time, delete_time, is_active FROM tokens
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetTokenForUpdate(ctx context.Context, id int64) (Token, error) {
	row := q.db.QueryRowContext(ctx, getTokenForUpdate, id)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.UID,
		&i.TokenName,
		&i.Symbol,
		&i.Supply,
		&i.ContractAddress,
		&i.CreateTime,
		&i.UpdateTime,
		&i.DeleteTime,
		&i.IsActive,
	)
	return i, err
}

const listTokens = `-- name: ListTokens :many
SELECT id, u_id, token_name, symbol, supply, contract_address, create_time, update_time, delete_time, is_active FROM tokens
WHERE u_id = $1
ORDER BY id
`

func (q *Queries) ListTokens(ctx context.Context, uID int64) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, listTokens, uID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Token{}
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.ID,
			&i.UID,
			&i.TokenName,
			&i.Symbol,
			&i.Supply,
			&i.ContractAddress,
			&i.CreateTime,
			&i.UpdateTime,
			&i.DeleteTime,
			&i.IsActive,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateToken = `-- name: UpdateToken :one
UPDATE tokens 
SET contract_address = $2
WHERE id = $1
RETURNING id, u_id, token_name, symbol, supply, contract_address, create_time, update_time, delete_time, is_active
`

type UpdateTokenParams struct {
	ID              int64  `json:"id"`
	ContractAddress string `json:"contract_address"`
}

func (q *Queries) UpdateToken(ctx context.Context, arg UpdateTokenParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, updateToken, arg.ID, arg.ContractAddress)
	var i Token
	err := row.Scan(
		&i.ID,
		&i.UID,
		&i.TokenName,
		&i.Symbol,
		&i.Supply,
		&i.ContractAddress,
		&i.CreateTime,
		&i.UpdateTime,
		&i.DeleteTime,
		&i.IsActive,
	)
	return i, err
}
