package main

import (
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/chacha20poly1305"
)

func main() {
	key := make([]byte, chacha20poly1305.KeySize) // 32 bytes
	rand.Read(key)

	aead, err := chacha20poly1305.New(key)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, aead.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	ciphertext := aead.Seal(nonce, nonce, []byte("hello ChaCha20-Poly1305"), nil)
	fmt.Printf("ChaCha20-Poly1305 ciphertext: %x\n", ciphertext)
}
