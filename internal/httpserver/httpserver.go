package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gustialfian/transfer-system-golang/internal/modules/account"
)

type HttpServerOpt struct {
	Addr    string
	Handler *ServiceHandler
}

type ServiceHandler struct {
	Account AccountHandler
}

type AccountHandler interface {
	Create(ctx context.Context, data account.AccountCreate) error
	ById(ctx context.Context, accountId int) (account.Account, error)
}

func NewMux(opt HttpServerOpt) *http.Server {
	h := opt.Handler
	r := http.NewServeMux()

	r.HandleFunc("POST /accounts", h.accountCreate)
	r.HandleFunc("GET /accounts/{account_id}", h.accountById)

	return &http.Server{
		Addr:              opt.Addr,
		Handler:           r,
		ReadHeaderTimeout: 1 * time.Second,
	}
}

func (h *ServiceHandler) accountCreate(w http.ResponseWriter, r *http.Request) {
}

func (h *ServiceHandler) accountById(w http.ResponseWriter, r *http.Request) {
}
