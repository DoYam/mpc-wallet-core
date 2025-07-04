package core

import "math/big"

type KeyShare struct {
	ID string
	D  *big.Int
	X  *big.Int
	Y  *big.Int
}

type PartialSignature struct {
	R *big.Int `json:"r"`
	S *big.Int `json:"s"`
}