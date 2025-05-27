package tigerbeetledb

import (
	"fmt"
	"log"

	tb "github.com/tigerbeetle/tigerbeetle-go"
	tbt "github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

func MustNewTigerbeetle(address string) *TigerBeetleDB {
	client, err := tb.NewClient(tbt.ToUint128(0), []string{address})
	if err != nil {
		log.Fatalf("tigerbeetle: error creating client: %v", err)
		return nil
	}

	return &TigerBeetleDB{client: client}
}

type TigerBeetleDB struct {
	client tb.Client
}

func (tdb *TigerBeetleDB) CreateAccount(accountId int) error {
	res, err := tdb.client.CreateAccounts([]tbt.Account{{
		ID:          tbt.ToUint128(uint64(accountId)),
		UserData128: tbt.ToUint128(uint64(accountId)),
		Ledger:      1,
		Code:        1,
	}})
	if err != nil {
		return fmt.Errorf("error creating accounts: %s", err)
	}
	for _, err := range res {
		return fmt.Errorf("error creating account %d: %s", err.Index, err.Result)
	}
	return nil
}

func (tdb *TigerBeetleDB) CreateTransaction(debitAccountId int, creditAccountId int, amount int) error {
	transferRes, err := tdb.client.CreateTransfers([]tbt.Transfer{
		{
			ID:              tbt.ID(),
			DebitAccountID:  tbt.ToUint128(uint64(debitAccountId)),
			CreditAccountID: tbt.ToUint128(uint64(creditAccountId)),
			Amount:          tbt.ToUint128(uint64(amount)),
			Ledger:          1,
			Code:            1,
		},
	})
	if err != nil {
		return fmt.Errorf("error creating transfer: %s", err)
	}
	for _, err := range transferRes {
		return fmt.Errorf("error creating transfer: %s", err.Result)
	}
	return nil
}

func (tdb *TigerBeetleDB) Close() {
	tdb.client.Close()
}
