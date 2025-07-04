package internal

import (
	"mpc-wallet-core/core"
)

var MyKeyShare *core.KeyShare

func InitKeyShare(id string) error {
	ksA, ksB, err := core.GenerateKeyShares()
	if err != nil {
		return err
	}

	if id == "A" {
		MyKeyShare = ksA
	} else {
		MyKeyShare = ksB
	}
	return nil

}

func Sign(msg string) (*core.PartialSignature, error) { return core.PartialSign(msg, MyKeyShare) }

func GetPublicKeyJSON() ([]byte, error) { return core.PublicKeyToJSON(MyKeyShare) }
