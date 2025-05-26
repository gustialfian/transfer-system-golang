package httpserver

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gustialfian/transfer-system-golang/internal/modules/account"
)

type AccountHandler interface {
	Create(ctx context.Context, data account.AccountCreate) error
	ById(ctx context.Context, accountId int) (account.Account, error)
}

func (h *ServiceHandler) accountCreate(w http.ResponseWriter, r *http.Request) {
	var body account.AccountCreate
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.Account.Create(r.Context(), body)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *ServiceHandler) accountById(w http.ResponseWriter, r *http.Request) {
	accountIdStr := r.PathValue("account_id")
	accountId, err := strconv.Atoi(accountIdStr)
	if err != nil {
		http.Error(w, "Invalid account_id", http.StatusBadRequest)
		return
	}

	data, err := h.Account.ById(r.Context(), accountId)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
