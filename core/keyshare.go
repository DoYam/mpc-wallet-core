package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

func GenerateKeyShares() (*KeyShare, *KeyShare, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	D1, _ := rand.Int(rand.Reader, priv.D)
	D2 := new(big.Int).Sub(priv.D, D1)

	pub := priv.PublicKey
	return &KeyShare{"A", D1, pub.X, pub.Y}, &KeyShare{"B", D2, pub.X, pub.Y}, nil

}
