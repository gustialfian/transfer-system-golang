package account

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// AccountDB provides methods for interacting with the accounts table in the database.
type AccountDB struct {
	db *sqlx.DB
}

// NewAccountDB creates and returns a new instance of AccountDB
func NewAccountDB(db *sqlx.DB) *AccountDB {
	return &AccountDB{db}
}

// AccountCreateParams holds the parameters required to create a new account.
type AccountCreateParams struct {
	AccountId    int
	Balance      int
	ScaleBalance int
}

// Create inserts a new account record into the accounts table with the provided parameters.
func (db *AccountDB) Create(ctx context.Context, params AccountCreateParams) error {
	q := `
	INSERT INTO accounts (account_id, balance, scale_balance, created_at, updated_at)
	VALUES ($1, $2, $3, NOW(), NOW())`
	if _, err := db.db.ExecContext(ctx, q, params.AccountId, params.Balance, params.ScaleBalance); err != nil {
		return fmt.Errorf("sql insert: %w [query: %s]", err, q)
	}

	return nil
}

// AccountRow represents a row in the accounts table, containing the account ID,
// the current balance, and the scaled balance for precision handling.
type AccountRow struct {
	AccountId    int `db:"account_id"`
	Balance      int `db:"balance"`
	ScaleBalance int `db:"scale_balance"`
}

// ById retrieves an account record from the database by its account ID.
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

// AccountUpdateBalanceParams contains the parameters required to update the balance of an account.
// It includes the unique identifier of the account and the new balance value to be set.
type AccountUpdateBalanceParams struct {
	AccountId int
	Balance   int
}

// UpdateBalance updates the balance of an account identified by AccountId in the database.
// It sets the new balance and updates the updated_at timestamp to the current time.
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
