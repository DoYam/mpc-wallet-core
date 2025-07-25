package handlers

import (
	"net/http"

	"mpc-wallet-core/internal"
)

func WalletHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		pub, err := internal.GetPublicKeyJSON()
		if err != nil {
			http.Error(w, "failed to get public key", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(pub)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}
