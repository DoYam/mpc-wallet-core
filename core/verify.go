package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/json"
	"math/big"
)

func Verify(msg string, r, s, x, y *big.Int) bool {
	hash := sha256.Sum256([]byte(msg))
	pub := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}
	return ecdsa.Verify(&pub, hash[:], r, s)
}

func PublicKeyToJSON(ks *KeyShare) ([]byte, error) {
	res := map[string]string{
		"curve": "P-256",
		"x":     ks.X.Text(16),
		"y":     ks.Y.Text(16),
	}
	return json.Marshal(res)
}