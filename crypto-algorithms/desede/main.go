package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"fmt"
)

func main() {
	key := []byte("24byte-triple-des-key!!!") // 3DES key = 24 bytes
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}

	plaintext := []byte("3desblok") // 8 bytes = one block
	iv := make([]byte, des.BlockSize)
	rand.Read(iv)

	ciphertext := make([]byte, len(plaintext))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext, plaintext)
	fmt.Printf("3DES (DESede) ciphertext: %x\n", ciphertext)
}
