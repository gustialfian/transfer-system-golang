// Package httpserver provides HTTP server setup and routing for the transfer system.
// It defines the ServiceHandler struct to aggregate account and transaction handlers,
// and exposes functions to create and configure an HTTP server with predefined routes
// for account and transaction operations.
package httpserver

import (
	"encoding/json"
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

// appResponse represents a standard HTTP JSON response structure.
type appResponse struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

// writeJSON writes the given appResponse as a JSON-encoded HTTP response with the specified status code.
func writeJSON(w http.ResponseWriter, statusCode int, resp appResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, `{"error":"failed to encode response"}`, http.StatusInternalServerError)
	}
}
