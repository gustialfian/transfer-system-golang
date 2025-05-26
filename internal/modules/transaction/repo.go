package transaction

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TransactionDB struct {
	db *sqlx.DB
}

type TransactionCreateParams struct {
	SourceAccountId      int
	DestinationAccountId int
	Amount               int
	AmountScale          int
}

func NewTransactionDB(db *sqlx.DB) *TransactionDB {
	return &TransactionDB{db}
}

func (db *TransactionDB) Create(ctx context.Context, params TransactionCreateParams) error {
	q := `
	INSERT INTO transactions (source_account_id, destination_account_id, amount, scale_amount, created_at, updated_at)
	VALUES ($1, $2, $3, $4, NOW(), NOW())`
	if _, err := db.db.ExecContext(ctx, q, params.SourceAccountId, params.DestinationAccountId, params.Amount, params.AmountScale); err != nil {
		return fmt.Errorf("sql insert: %w [query: %s]", err, q)
	}

	return nil
}
