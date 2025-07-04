package handlers

import (
	"encoding/json"
	"net/http"
	"mpc-wallet-core/internal"
)

type VerifyRequest struct {
	Message string `json:"message"`
	R       string `json:"r"`
	S       string `json:"s"`
}

type VerifyResponse struct {
	Valid bool `json:"valid"`
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	isValid := internal.VerifySignature(req.Message, req.R, req.S)

	resp := VerifyResponse{Valid: isValid}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
