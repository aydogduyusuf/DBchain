package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db: db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	TransactionType string 	`json:transaction_type`
	FromAddress 	string 	`json:"from_address"`
	ToAddress   	string 	`json:"to_address"`
	Data        	string 	`json:"data"`
	HashValue		string	`json:"hash_value"`
}

type TransferTxResult struct {
	Transaction	Transaction `json:"Transfer"`
}

func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	
	err := store.execTx(ctx, func(q *Queries) error  { 
		var err error

		result.Transaction, err = q.CreateTransaction(ctx, CreateTransactionParams{
			TransactionType: arg.TransactionType,
			FromAddress: arg.FromAddress,
			ToAddress: arg.ToAddress,
			TransferData: arg.Data,
			HashValue: arg.HashValue,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
