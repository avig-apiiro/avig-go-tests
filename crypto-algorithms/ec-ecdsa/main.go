package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256([]byte("message to sign"))
	sig, err := ecdsa.SignASN1(rand.Reader, priv, hash[:])
	if err != nil {
		panic(err)
	}

	valid := ecdsa.VerifyASN1(&priv.PublicKey, hash[:], sig)
	fmt.Printf("ECDSA signature: %x\nvalid: %v\n", sig, valid)
}
