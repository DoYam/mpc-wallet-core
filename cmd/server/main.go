package main

import (
	"fmt"
	"log"
	"net/http"

	"mpc-wallet-core/handlers"
	"mpc-wallet-core/internal"
)

func main() {
	// Initialize key share for Party A (could be B on another server)
	if err := internal.InitKeyShare("A"); err != nil {
		log.Fatalf("failed to init key share: %v", err)
	}

	http.HandleFunc("/partial-sign", handlers.PartialSignHandler)
	http.HandleFunc("/wallet", handlers.WalletHandler)
	http.HandleFunc("/combine", handlers.CombineHandler)

	fmt.Println("[MPC Node A] Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
