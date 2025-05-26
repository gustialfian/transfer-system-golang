package transaction

import (
	"context"
	"errors"
	"log"

	"github.com/gustialfian/transfer-system-golang/internal/modules/money"
)

type TransactionService struct {
	repo TransactionRepo
}

type TransactionRepo interface {
	Create(ctx context.Context, data TransactionCreateParams) error
}

type TransactionCreate struct {
	SourceAccountId      int    `json:"source_account_id"`
	DestinationAccountId int    `json:"destination_account_id"`
	Amount               string `json:"amount"`
}

type Transaction struct {
	SourceAccountId      int    `json:"source_account_id"`
	DestinationAccountId int    `json:"destination_account_id"`
	Amount               string `json:"amount"`
}

var ErrTransactionCreateFailed = errors.New("transaction creation fail")

func NewAccountService(repo TransactionRepo) *TransactionService {
	return &TransactionService{repo}
}

func (svc *TransactionService) Create(ctx context.Context, data TransactionCreate) error {
	var params TransactionCreateParams
	params.SourceAccountId = data.SourceAccountId
	params.DestinationAccountId = data.DestinationAccountId

	amount, err := money.StringToInt(data.Amount, money.Scale)
	if err != nil {
		log.Printf("%s: %s\n", ErrTransactionCreateFailed, err)
		return ErrTransactionCreateFailed
	}
	params.Amount = amount
	params.AmountScale = money.Scale

	if err := svc.repo.Create(ctx, params); err != nil {
		log.Printf("%s: %s\n", ErrTransactionCreateFailed, err)
		return ErrTransactionCreateFailed
	}
	return nil
}
