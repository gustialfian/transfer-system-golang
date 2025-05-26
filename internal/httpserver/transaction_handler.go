package httpserver

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gustialfian/transfer-system-golang/internal/modules/transaction"
)

type TransactionHandler interface {
	Create(ctx context.Context, data transaction.TransactionCreate) error
}

func (h *ServiceHandler) transactionCreate(w http.ResponseWriter, r *http.Request) {
	var body transaction.TransactionCreate
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.Transaction.Create(r.Context(), body)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
