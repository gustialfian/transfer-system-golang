package db

import (
	"context"
	"fmt"

	"github.com/gustialfian/transfer-system-golang/internal/domains/transaction"
	"github.com/jmoiron/sqlx"
)

// TransactionDB provides methods for interacting with the transactions table in the database.
type TransactionDB struct {
	db *sqlx.DB
}

// NewAccountDB creates and returns a new instance of TransactionDB
func NewTransactionDB(db *sqlx.DB) *TransactionDB {
	return &TransactionDB{db}
}

// Create inserts a new transaction record into the transactions table with the provided parameters.
func (db *TransactionDB) Create(ctx context.Context, params transaction.TransactionCreateParams) error {
	q := `
	INSERT INTO transactions (source_account_id, destination_account_id, amount, scale_amount, created_at, updated_at)
	VALUES ($1, $2, $3, $4, NOW(), NOW())`
	if _, err := db.db.ExecContext(ctx, q, params.SourceAccountId, params.DestinationAccountId, params.Amount, params.AmountScale); err != nil {
		return fmt.Errorf("sql insert: %w [query: %s]", err, q)
	}

	return nil
}
