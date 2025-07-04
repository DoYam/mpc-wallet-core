package handlers

import (
	"encoding/json"
	"net/http"

	"mpc-wallet-core/core"
)

type CombineRequest struct {
	SigA core.PartialSignature `json:"sig_a"`
	SigB core.PartialSignature `json:"sig_b"`
}

func CombineHandler(w http.ResponseWriter, r *http.Request) {
	var req CombineRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	combined := core.CombineSignatures(&req.SigA, &req.SigB)

	resp := map[string]string{
		"r": combined.R.Text(16),
		"s": combined.S.Text(16),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}