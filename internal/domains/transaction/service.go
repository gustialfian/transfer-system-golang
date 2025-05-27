package transaction

import (
	"context"
	"errors"
	"log"

	"github.com/gustialfian/transfer-system-golang/internal/domains/account"
	"github.com/gustialfian/transfer-system-golang/internal/domains/money"
)

// TransactionService provides methods for managing transactions.
type TransactionService struct {
	repo        TransactionRepo
	accountRepo account.AccountRepo
}

var (
	ErrTransactionCreateFailed               = errors.New("transaction creation fail")
	ErrTransactionSourceAccountNotFound      = errors.New("transaction source account not found")
	ErrTransactionDestinationAccountNotFound = errors.New("transaction destination account not found")
	ErrTransactionSourceBalanceNotEnough     = errors.New("transaction source balance not enough")
	ErrTransactionSourceBalanceNegative      = errors.New("transaction source balance negative")
	ErrTransactionSourceDestinationSame      = errors.New("transaction source and destination account can not be the same")
)

// NewTransactionService creates a new TransactionService with the given dependency.
func NewTransactionService(repo TransactionRepo, accountRepo account.AccountRepo) *TransactionService {
	return &TransactionService{repo, accountRepo}
}

// TransactionCreate represents the required information to create a new transaction
type TransactionCreate struct {
	SourceAccountId      int    `json:"source_account_id"`
	DestinationAccountId int    `json:"destination_account_id"`
	Amount               string `json:"amount"`
}

// Create executes a transaction by validating input, checking balances, updating accounts, and recording the transaction.
func (svc *TransactionService) Create(ctx context.Context, data TransactionCreate) error {
	amount, err := money.StringToInt(data.Amount, money.Scale)
	if err != nil {
		log.Printf("%s: %s\n", money.ErrMoneyParseFail, err)
		return money.ErrMoneyParseFail
	}

	if amount < 0 {
		log.Printf("%s\n", ErrTransactionSourceBalanceNegative)
		return ErrTransactionSourceBalanceNegative
	}

	if data.SourceAccountId == data.DestinationAccountId {
		log.Printf("%s\n", ErrTransactionSourceDestinationSame)
		return ErrTransactionSourceDestinationSame
	}

	destinationAccount, err := svc.accountRepo.ById(ctx, data.DestinationAccountId)
	if err != nil {
		log.Printf("%s: %s\n", ErrTransactionDestinationAccountNotFound, err)
		return ErrTransactionDestinationAccountNotFound
	}

	sourceAccount, err := svc.accountRepo.ById(ctx, data.SourceAccountId)
	if err != nil {
		log.Printf("%s: %s\n", ErrTransactionSourceAccountNotFound, err)
		return ErrTransactionSourceAccountNotFound
	}

	destinationBalance := destinationAccount.Balance + amount
	sourceBalance := sourceAccount.Balance - amount
	if sourceBalance < 0 {
		log.Printf("%s\n", ErrTransactionSourceBalanceNotEnough)
		return ErrTransactionSourceBalanceNotEnough
	}

	err = svc.accountRepo.UpdateBalance(ctx, account.AccountUpdateBalanceParams{
		AccountId: data.SourceAccountId,
		Balance:   sourceBalance,
	})
	if err != nil {
		log.Printf("%s: %s\n", ErrTransactionCreateFailed, err)
		return ErrTransactionCreateFailed
	}

	err = svc.accountRepo.UpdateBalance(ctx, account.AccountUpdateBalanceParams{
		AccountId: data.DestinationAccountId,
		Balance:   destinationBalance,
	})
	if err != nil {
		log.Printf("%s: %s\n", ErrTransactionCreateFailed, err)
		return ErrTransactionCreateFailed
	}

	params := TransactionCreateParams{
		SourceAccountId:      data.SourceAccountId,
		DestinationAccountId: data.DestinationAccountId,
		Amount:               amount,
		AmountScale:          money.Scale,
	}

	if err := svc.repo.Create(ctx, params); err != nil {
		log.Printf("%s: %s\n", ErrTransactionCreateFailed, err)
		return ErrTransactionCreateFailed
	}

	return nil
}
