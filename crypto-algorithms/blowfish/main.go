package main

import (
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/blowfish"
)

func main() {
	key := []byte("a-blowfish-key")
	block, err := blowfish.NewCipher(key)
	if err != nil {
		panic(err)
	}

	plaintext := []byte("blowfish8") // must be padded to block size (8)
	plaintext = append(plaintext, make([]byte, blowfish.BlockSize-len(plaintext)%blowfish.BlockSize)...)

	iv := make([]byte, blowfish.BlockSize)
	rand.Read(iv)

	ciphertext := make([]byte, len(plaintext))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext, plaintext)
	fmt.Printf("Blowfish ciphertext: %x\n", ciphertext)
}
