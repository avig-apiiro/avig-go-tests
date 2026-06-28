package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"fmt"
)

func main() {
	key := []byte("8bytekey") // DES key = 8 bytes
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	plaintext := []byte("desblock") // 8 bytes = one block
	iv := make([]byte, des.BlockSize)
	rand.Read(iv)

	ciphertext := make([]byte, len(plaintext))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext, plaintext)
	fmt.Printf("DES ciphertext: %x\n", ciphertext)
}
