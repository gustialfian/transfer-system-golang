package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gustialfian/transfer-system-golang/internal/domains/transaction"
)

// TransactionHandler is interface that ServiceHandler use to integrate with TransactionService
type TransactionHandler interface {
	Create(ctx context.Context, data transaction.TransactionCreate) error
}

func (h *ServiceHandler) transactionCreate(w http.ResponseWriter, r *http.Request) {
	var body transaction.TransactionCreate
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeJSON(w, http.StatusBadRequest, appResponse{Error: "bad request"})
		return
	}

	err := h.Transaction.Create(r.Context(), body)
	if err != nil {
		if errors.Is(err, transaction.ErrTransactionSourceAccountNotFound) {
			writeJSON(w, http.StatusNotFound, appResponse{Error: "transaction source account not found"})

		} else if errors.Is(err, transaction.ErrTransactionDestinationAccountNotFound) {
			writeJSON(w, http.StatusNotFound, appResponse{Error: "transaction destination account not found"})

		} else if errors.Is(err, transaction.ErrTransactionSourceBalanceNotEnough) {
			writeJSON(w, http.StatusBadRequest, appResponse{Error: "transaction source balance not enough"})

		} else if errors.Is(err, transaction.ErrTransactionSourceDestinationSame) {
			writeJSON(w, http.StatusBadRequest, appResponse{Error: "transaction source and destination account can not be the same"})

		} else {
			writeJSON(w, http.StatusBadRequest, appResponse{Error: "bad request"})
		}
		return
	}

	writeJSON(w, http.StatusOK, appResponse{Message: "transaction created"})
}
