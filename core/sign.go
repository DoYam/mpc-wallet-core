package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
)

func PartialSign(msg string, ks *KeyShare) (*PartialSignature, error) {
	hash := sha256.Sum256([]byte(msg))

	priv := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     ks.X,
			Y:     ks.Y,
		},
		D: ks.D,
	}

	r, s, err := ecdsa.Sign(rand.Reader, priv, hash[:])
	if err != nil {
		return nil, err
	}

	return &PartialSignature{R: r, S: s}, nil

}
