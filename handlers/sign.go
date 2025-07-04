package handlers

import (
	"encoding/json"
	"net/http"

	"mpc-wallet-core/internal"
)

type SignRequest struct {
	Message string `json:"message"`
}

type SignResponse struct {
	R         string `json:"r"`
	S         string `json:"s"`
	Signature string `json:"signature"`
}

func SignHandler(w http.ResponseWriter, r *http.Request) {
	var req SignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	rVal, sVal, sig, err := internal.SignMessage(req.Message)
	if err != nil {
		http.Error(w, "failed to sign", http.StatusInternalServerError)
		return
	}

	resp := SignResponse{R: rVal, S: sVal, Signature: sig}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
