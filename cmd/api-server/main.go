package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gustialfian/transfer-system-golang/internal/config"
	"github.com/gustialfian/transfer-system-golang/internal/db"
	"github.com/gustialfian/transfer-system-golang/internal/httpserver"
	"github.com/gustialfian/transfer-system-golang/internal/modules/account"
	"github.com/gustialfian/transfer-system-golang/internal/modules/transaction"
)

func main() {
	cfg := config.LoadConfig()

	db := db.MustNewPostgreSQL(cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresDBName)

	accountRepo := account.NewAccountDB(db)
	accountSvc := account.NewAccountService(accountRepo)

	transactionRepo := transaction.NewTransactionDB(db)
	transactionSvc := transaction.NewAccountService(transactionRepo, accountRepo)

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
