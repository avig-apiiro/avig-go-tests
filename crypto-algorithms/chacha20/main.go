package main

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/chacha20"
)

func main() {
	key := make([]byte, chacha20.KeySize)     // 32 bytes
	nonce := make([]byte, chacha20.NonceSize) // 12 bytes
	rand.Read(key)
	rand.Read(nonce)

	c, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		panic(err)
	}

	plaintext := []byte("hello ChaCha20")
	ciphertext := make([]byte, len(plaintext))
	c.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("ChaCha20 ciphertext: %x\n", ciphertext)
}
