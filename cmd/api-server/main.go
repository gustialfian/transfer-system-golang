package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gustialfian/transfer-system-golang/internal/config"
	"github.com/gustialfian/transfer-system-golang/internal/db"
	"github.com/gustialfian/transfer-system-golang/internal/httpserver"
	"github.com/gustialfian/transfer-system-golang/internal/modules/account"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.LoadConfig()

	db := db.MustNewPostgreSQL(cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresDBName)

	handler := newServiceHandler(db)

	server := httpserver.NewMux(httpserver.HttpServerOpt{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: handler,
	})

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

func newServiceHandler(db *sqlx.DB) *httpserver.ServiceHandler {
	handler := &httpserver.ServiceHandler{}

	accountRepo := account.NewAccountDB(db)
	handler.Account = account.NewAccountService(accountRepo)

	return handler
}
