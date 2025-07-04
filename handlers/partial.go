package handlers

import (
	"encoding/json"
	"net/http"

	"mpc-wallet-core/internal"
)

type SignRequest struct {
	Message string `json:"message"`
}

func PartialSignHandler(w http.ResponseWriter, r *http.Request) {
	var req SignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	sig, err := internal.Sign(req.Message)
	if err != nil {
		http.Error(w, "failed to sign", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sig)
}