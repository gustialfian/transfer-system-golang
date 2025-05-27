package account

import "context"

// AccountRepo defines the interface for account data persistence.
// Implementations of this interface handle the actual data storage and retrieval.
type AccountRepo interface {
	Create(ctx context.Context, data AccountCreateParams) error
	ById(ctx context.Context, accountId int) (AccountRow, error)
	UpdateBalance(ctx context.Context, params AccountUpdateBalanceParams) error
}

// AccountCreateParams holds the parameters required to create a new account.
type AccountCreateParams struct {
	AccountId    int
	Balance      int
	ScaleBalance int
}

// AccountRow represents a row in the accounts table, containing the account ID,
// the current balance, and the scaled balance for precision handling.
type AccountRow struct {
	AccountId    int `db:"account_id"`
	Balance      int `db:"balance"`
	ScaleBalance int `db:"scale_balance"`
}

// AccountUpdateBalanceParams contains the parameters required to update the balance of an account.
// It includes the unique identifier of the account and the new balance value to be set.
type AccountUpdateBalanceParams struct {
	AccountId int
	Balance   int
}
