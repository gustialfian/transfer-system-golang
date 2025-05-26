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

func (db *AccountDB) Create(ctx context.Context, params AccountCreate) error {
	return fmt.Errorf("not impl")
}

func (db *AccountDB) byId(ctx context.Context, accountId int) (Account, error) {
	return Account{}, fmt.Errorf("not impl")
}
