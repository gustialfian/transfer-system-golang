// Package account provides business logic for interacting with accounts,
// including creation, retrieval, and balance management.
package account

import (
	"context"
	"errors"
	"log"

	"github.com/gustialfian/transfer-system-golang/internal/domains/money"
)

// AccountService encapsulates account-related operations and business logic.
type AccountService struct {
	repo AccountRepo
}

// AccountCreate represents the parameters required to create a new account.
type AccountCreate struct {
	AccountId      int    `json:"account_id"`      // Unique identifier for the account.
	InitialBalance string `json:"initial_balance"` // Initial balance as a string (e.g., "100.00").
}

// Account represents an account with its ID and balance.
type Account struct {
	AccountId      int    `json:"account_id"`      // Unique identifier for the account.
	InitialBalance string `json:"initial_balance"` // Balance as a string (e.g., "100.00").
}

var (
	ErrAccountCreateFailed           = errors.New("account creation fail")
	ErrAccountByIdFailed             = errors.New("account by id fail")
	ErrAccountInitialBalanceNegative = errors.New("account initial balance negative")
)

// NewAccountService creates a new AccountService with the given repository.
func NewAccountService(repo AccountRepo) *AccountService {
	return &AccountService{repo}
}

// Create creates a new account with the specified initial balance.
func (svc *AccountService) Create(ctx context.Context, data AccountCreate) error {
	initialBalance, err := money.StringToInt(data.InitialBalance, money.Scale)
	if err != nil {
		log.Printf("%s: %s\n", ErrAccountCreateFailed, err)
		return money.ErrMoneyParseFail
	}

	if initialBalance < 0 {
		log.Printf("%s\n", ErrAccountInitialBalanceNegative)
		return ErrAccountInitialBalanceNegative
	}

	params := AccountCreateParams{
		AccountId:    data.AccountId,
		Balance:      initialBalance,
		ScaleBalance: money.Scale,
	}

	if err := svc.repo.Create(ctx, params); err != nil {
		log.Printf("%s: %s\n", ErrAccountCreateFailed, err)
		return ErrAccountCreateFailed
	}
	return nil
}

// ById retrieves an account by its ID.
func (svc *AccountService) ById(ctx context.Context, accountId int) (Account, error) {
	row, err := svc.repo.ById(ctx, accountId)
	if err != nil {
		log.Printf("%s: %s\n", ErrAccountByIdFailed, err)
		return Account{}, ErrAccountByIdFailed
	}

	initialBalance := money.IntToString(row.Balance, row.ScaleBalance)
	data := Account{
		AccountId:      row.AccountId,
		InitialBalance: initialBalance,
	}

	return data, nil
}
