package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gustialfian/transfer-system-golang/internal/modules/transaction"
)

// TransactionHandler is interface that ServiceHandler use to integrate with TransactionService
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
		if errors.Is(err, transaction.ErrTransactionSourceAccountNotFound) {
			http.Error(w, "transaction source account not found", http.StatusBadRequest)

		} else if errors.Is(err, transaction.ErrTransactionDestinationAccountNotFound) {
			http.Error(w, "transaction destination account not found", http.StatusBadRequest)

		} else if errors.Is(err, transaction.ErrTransactionSourceBalanceNotEnough) {
			http.Error(w, "transaction source balance not enough", http.StatusBadRequest)

		} else if errors.Is(err, transaction.ErrTransactionSourceDestinationSame) {
			http.Error(w, "transaction source and destination account can not be the same", http.StatusBadRequest)

		} else {
			http.Error(w, "bad request", http.StatusBadRequest)
		}
		return
	}
}
