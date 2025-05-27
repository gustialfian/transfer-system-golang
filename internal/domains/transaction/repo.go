package transaction

import (
	"context"
)

// TransactionRepo defines the interface for transaction repository operations.
type TransactionRepo interface {
	Create(ctx context.Context, data TransactionCreateParams) error
}

// AccountCreateParams holds the parameters required to create a new transaction.
type TransactionCreateParams struct {
	SourceAccountId      int
	DestinationAccountId int
	Amount               int
	AmountScale          int
}

type TransactionTBRepo interface {
	CreateTransaction(debitAccountId int, creditAccountId int, amount int) error
}
