package httpserver

import (
	"net/http"
	"time"
)

type HttpServerOpt struct {
	Addr    string
	Handler *ServiceHandler
}

type ServiceHandler struct {
	Account     AccountHandler
	Transaction TransactionHandler
}

func NewMux(opt HttpServerOpt) *http.Server {
	h := opt.Handler
	r := http.NewServeMux()

	r.HandleFunc("POST /accounts", h.accountCreate)
	r.HandleFunc("GET /accounts/{account_id}", h.accountById)
	r.HandleFunc("POST /transactions", h.transactionCreate)

	return &http.Server{
		Addr:              opt.Addr,
		Handler:           r,
		ReadHeaderTimeout: 1 * time.Second,
	}
}
