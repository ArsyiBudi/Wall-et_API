package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all function to execute db queries and transaction
type Store struct {
	*Queries
	db *sql.DB
}

// newStore creates a NewStore
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
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

// TransferTxParams contains the input parameters fo the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer      Transfer `json:"transfer"`
	FromAccountID Account  `json:"from_account"`
	ToAccountID   Account  `json:"to_account"`
	FromEntry     Entry    `json:"from_entry"`
	ToEntry       Entry    `json:"to_entry"`
}

// transferTX performs a money transfer from one account to the other
// It creates a transfer record, add account entries, and update account balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	// create and running new transaction
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return nil
		}

		result.ToEntry, err = q.CreateEntries(ctx, CreateEntriesParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return nil
		}

		return nil

		// TODO: update account' balance

	})

	return result, err
}
