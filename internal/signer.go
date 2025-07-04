package internal

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

var privKey *ecdsa.PrivateKey

func InitKey() error {
	var err error
	privKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}
	fmt.Println("[Shared Key] Public:", privKey.PublicKey)
	return nil
}

func RefreshKey() error {
	return InitKey()
}

func GetPublicKeyJSON() ([]byte, error) {
	pub := privKey.PublicKey
	res := map[string]string{
		"curve": "P-256",
		"x":     pub.X.Text(16),
		"y":     pub.Y.Text(16),
	}
	return json.Marshal(res)
}

func SignMessage(msg string) (string, string, string, error) {
	hash := sha256.Sum256([]byte(msg))
	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
	if err != nil {
		return "", "", "", err
	}
	return r.Text(16), s.Text(16), fmt.Sprintf("%x|%x", r, s), nil
}