package account

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AccountDB struct {
	db *sqlx.DB
}

func NewAccountDB(db *sqlx.DB) *AccountDB {
	return &AccountDB{db}
}

type AccountCreateParams struct {
	AccountId      int
	InitialBalance int
	ScaleBalance   int
}

func (db *AccountDB) Create(ctx context.Context, params AccountCreateParams) error {
	q := `
	INSERT INTO accounts (account_id, initial_balance, scale_balance, created_at, updated_at)
	VALUES ($1, $2, $3, NOW(), NOW())`
	if _, err := db.db.ExecContext(ctx, q, params.AccountId, params.InitialBalance, params.ScaleBalance); err != nil {
		return fmt.Errorf("sql insert: %w [query: %s]", err, q)
	}

	return nil
}

type AccountRow struct {
	AccountId      int `db:"account_id"`
	InitialBalance int `db:"initial_balance"`
	ScaleBalance   int `db:"scale_balance"`
}

func (db *AccountDB) ById(ctx context.Context, accountId int) (AccountRow, error) {
	var rows []AccountRow

	q := `
	SELECT x.account_id
		, x.initial_balance
		, x.scale_balance
	FROM accounts AS x
	WHERE x.account_id = $1`
	err := sqlx.SelectContext(ctx, db.db, &rows, q, accountId)
	if err != nil {
		return AccountRow{}, fmt.Errorf("sql select: %w [query: %s]", err, q)
	}

	if len(rows) == 0 {
		return AccountRow{}, fmt.Errorf("account not found [account_id: %d]", accountId)
	}

	return rows[0], nil
}
