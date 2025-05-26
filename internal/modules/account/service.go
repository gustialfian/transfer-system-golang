package account

import (
	"context"
	"errors"
	"log"
)

type AccountService struct {
	repo AccountRepo
}

type AccountRepo interface {
	Create(ctx context.Context, data AccountCreate) error
	byId(ctx context.Context, accountId int) (Account, error)
}

type AccountCreate struct {
	AccountId      int    `json:"account_id"`
	InitialBalance uint64 `json:"initial_balance"`
}

type Account struct {
	AccountId      int    `json:"account_id"`
	InitialBalance uint64 `json:"initial_balance"`
}

var (
	ErrAccountCreateFailed = errors.New("account creation fail")
	ErrAccountByIdFailed   = errors.New("account creation fail")
)

func NewAccountService(repo AccountRepo) *AccountService {
	return &AccountService{repo}
}

func (svc *AccountService) Create(ctx context.Context, data AccountCreate) error {
	if err := svc.repo.Create(ctx, data); err != nil {
		log.Printf("%s: %s\n", ErrAccountCreateFailed, err)
		return ErrAccountCreateFailed
	}
	return nil
}

func (svc *AccountService) ById(ctx context.Context, accountId int) (Account, error) {
	data, err := svc.repo.byId(ctx, accountId)
	if err != nil {
		log.Printf("%s: %s\n", ErrAccountByIdFailed, err)
		return data, ErrAccountByIdFailed
	}

	return data, nil
}
