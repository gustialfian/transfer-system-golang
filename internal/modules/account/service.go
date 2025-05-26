package account

import (
	"context"
	"errors"
	"log"

	"github.com/gustialfian/transfer-system-golang/internal/modules/money"
)

type AccountService struct {
	repo AccountRepo
}

type AccountRepo interface {
	Create(ctx context.Context, data AccountCreateParams) error
	ById(ctx context.Context, accountId int) (AccountRow, error)
}

type AccountCreate struct {
	AccountId      int    `json:"account_id"`
	InitialBalance string `json:"initial_balance"`
}

type Account struct {
	AccountId      int    `json:"account_id"`
	InitialBalance string `json:"initial_balance"`
}

var (
	ErrAccountCreateFailed = errors.New("account creation fail")
	ErrAccountByIdFailed   = errors.New("account by id fail")
)

func NewAccountService(repo AccountRepo) *AccountService {
	return &AccountService{repo}
}

func (svc *AccountService) Create(ctx context.Context, data AccountCreate) error {
	var params AccountCreateParams
	params.AccountId = data.AccountId

	initialBalance, err := money.StringToInt(data.InitialBalance, money.Scale)
	if err != nil {
		log.Printf("%s: %s\n", ErrAccountCreateFailed, err)
		return ErrAccountCreateFailed
	}
	params.InitialBalance = initialBalance
	params.ScaleBalance = money.Scale

	if err := svc.repo.Create(ctx, params); err != nil {
		log.Printf("%s: %s\n", ErrAccountCreateFailed, err)
		return ErrAccountCreateFailed
	}
	return nil
}

func (svc *AccountService) ById(ctx context.Context, accountId int) (Account, error) {
	row, err := svc.repo.ById(ctx, accountId)
	if err != nil {
		log.Printf("%s: %s\n", ErrAccountByIdFailed, err)
		return Account{}, ErrAccountByIdFailed
	}

	initialBalance := money.IntToString(row.InitialBalance, row.ScaleBalance)
	data := Account{
		AccountId:      row.AccountId,
		InitialBalance: initialBalance,
	}

	return data, nil
}
