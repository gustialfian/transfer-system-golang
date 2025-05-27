package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gustialfian/transfer-system-golang/internal/config"
	"github.com/gustialfian/transfer-system-golang/internal/db"
	"github.com/gustialfian/transfer-system-golang/internal/domains/account"
	"github.com/gustialfian/transfer-system-golang/internal/domains/transaction"
	"github.com/gustialfian/transfer-system-golang/internal/httpserver"
)

func main() {
	cfg := config.LoadConfig()

	dbConn := db.MustNewPostgreSQL(cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresDBName)

	accountRepo := db.NewAccountDB(dbConn)
	accountSvc := account.NewAccountService(accountRepo)

	transactionRepo := db.NewTransactionDB(dbConn)
	transactionSvc := transaction.NewTransactionService(transactionRepo, accountRepo)

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
