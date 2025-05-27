package httpserver

import (
	"net/http"
	"time"
)

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

type ServiceHandler struct {
	Account     AccountHandler
	Transaction TransactionHandler
}
