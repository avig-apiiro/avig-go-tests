package main

import (
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
)

func main() {
	curve := ecdh.P256()

	alicePriv, _ := curve.GenerateKey(rand.Reader)
	bobPriv, _ := curve.GenerateKey(rand.Reader)

	aliceShared, _ := alicePriv.ECDH(bobPriv.PublicKey())
	bobShared, _ := bobPriv.ECDH(alicePriv.PublicKey())

	fmt.Printf("ECDH shared secret (alice): %x\n", aliceShared)
	fmt.Printf("ECDH shared secret (bob):   %x\n", bobShared)
}
