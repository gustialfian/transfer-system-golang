package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gustialfian/transfer-system-golang/internal/domains/account"
)

// AccountHandler is interface that ServiceHandler use to integrate with AccountService
type AccountHandler interface {
	Create(ctx context.Context, data account.AccountCreate) error
	ById(ctx context.Context, accountId int) (account.Account, error)
}

func (h *ServiceHandler) accountCreate(w http.ResponseWriter, r *http.Request) {
	var body account.AccountCreate
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeJSON(w, http.StatusBadRequest, appResponse{
			Error: "bad request",
		})
		return
	}

	err := h.Account.Create(r.Context(), body)
	if err != nil {
		if errors.Is(err, account.ErrAccountInitialBalanceNegative) {
			writeJSON(w, http.StatusBadRequest, appResponse{
				Error: "account initial balance is negative",
			})
		} else {
			writeJSON(w, http.StatusBadRequest, appResponse{
				Error: "bad request",
			})
		}
		return
	}

	writeJSON(w, http.StatusOK, appResponse{Message: "account created"})
}

func (h *ServiceHandler) accountById(w http.ResponseWriter, r *http.Request) {
	accountIdStr := r.PathValue("account_id")
	accountId, err := strconv.Atoi(accountIdStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, appResponse{
			Error: "invalid account_id",
		})
		return
	}

	data, err := h.Account.ById(r.Context(), accountId)
	if err != nil {
		writeJSON(w, http.StatusNotFound, appResponse{
			Error: "account not found",
		})
		return
	}

	writeJSON(w, http.StatusOK, appResponse{Data: data})
}
