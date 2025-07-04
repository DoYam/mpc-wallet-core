package main

import (
	"fmt"
	"log"
	"net/http"

	"mpc-wallet-core/handlers"
	"mpc-wallet-core/internal"
)

func main() {
	if err := internal.InitKey(); err != nil {
		log.Fatalf("failed to init key: %v", err)
	}

	http.HandleFunc("/sign", handlers.SignHandler)
	http.HandleFunc("/wallet", handlers.WalletHandler)
	http.HandleFunc("/verify", handlers.VerifyHandler)


	fmt.Println("[MPC PoC] Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}