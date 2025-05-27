package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gustialfian/transfer-system-golang/internal/domains/account"
	"github.com/gustialfian/transfer-system-golang/internal/domains/transaction"
	"github.com/gustialfian/transfer-system-golang/internal/infrastructure/config"
	"github.com/gustialfian/transfer-system-golang/internal/infrastructure/db"
	"github.com/gustialfian/transfer-system-golang/internal/infrastructure/httpserver"
	"github.com/gustialfian/transfer-system-golang/internal/infrastructure/tigerbeetledb"
)

func main() {
	cfg := config.LoadConfig()

	dbConn := db.MustNewPostgreSQL(cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresDBName)
	defer dbConn.Close()

	tigerbeetleDB := &tigerbeetledb.TigerBeetleDB{}
	if cfg.IsTigerBeetleOn {
		tigerbeetleDB = tigerbeetledb.MustNewTigerbeetle(cfg.TigerbeetleAddress)
	}
	defer tigerbeetleDB.Close()

	accountRepo := db.NewAccountDB(dbConn)
	accountSvc := account.NewAccountService(accountRepo, tigerbeetleDB, cfg.IsTigerBeetleOn)

	transactionRepo := db.NewTransactionDB(dbConn)
	transactionSvc := transaction.NewTransactionService(transactionRepo, accountRepo, tigerbeetleDB, cfg.IsTigerBeetleOn)

	handler := &httpserver.ServiceHandler{
		Account:     accountSvc,
		Transaction: transactionSvc,
	}

	server := httpserver.NewMux(fmt.Sprintf(":%s", cfg.Port), handler)

	log.Printf("listen on :%s\n", cfg.Port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}
