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
	AccountId    int
	Balance      int
	ScaleBalance int
}

func (db *AccountDB) Create(ctx context.Context, params AccountCreateParams) error {
	q := `
	INSERT INTO accounts (account_id, balance, scale_balance, created_at, updated_at)
	VALUES ($1, $2, $3, NOW(), NOW())`
	if _, err := db.db.ExecContext(ctx, q, params.AccountId, params.Balance, params.ScaleBalance); err != nil {
		return fmt.Errorf("sql insert: %w [query: %s]", err, q)
	}

	return nil
}

type AccountRow struct {
	AccountId    int `db:"account_id"`
	Balance      int `db:"balance"`
	ScaleBalance int `db:"scale_balance"`
}

func (db *AccountDB) ById(ctx context.Context, accountId int) (AccountRow, error) {
	var rows []AccountRow

	q := `
	SELECT x.account_id
		, x.balance
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

type AccountUpdateBalanceParams struct {
	AccountId int
	Balance   int
}

func (db *AccountDB) UpdateBalance(ctx context.Context, params AccountUpdateBalanceParams) error {
	q := `
	UPDATE accounts
	SET balance = $2
		, updated_at = NOW()
	WHERE account_id = $1`
	if _, err := db.db.ExecContext(ctx, q, params.AccountId, params.Balance); err != nil {
		return fmt.Errorf("sql insert: %w [query: %s]", err, q)
	}

	return nil
}
