// Package httpserver provides HTTP server setup and routing for the transfer system.
// It defines the ServiceHandler struct to aggregate account and transaction handlers,
// and exposes functions to create and configure an HTTP server with predefined routes
// for account and transaction operations.
package httpserver

import (
	"net/http"
	"time"
)

// NewMux creates and configures a new HTTP server with predefined routes
func NewMux(addr string, h *ServiceHandler) *http.Server {
	r := http.NewServeMux()

	r.HandleFunc("POST /accounts", h.accountCreate)
	r.HandleFunc("GET /accounts/{account_id}", h.accountById)
	r.HandleFunc("POST /transactions", h.transactionCreate)

	return &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: 1 * time.Second,
	}
}

// ServiceHandler aggregates handlers for account and transaction services,
// providing a unified interface for handling HTTP requests related to accounts
// and transactions within the system.
type ServiceHandler struct {
	Account     AccountHandler
	Transaction TransactionHandler
}
