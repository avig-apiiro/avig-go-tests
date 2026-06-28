package main

import (
	"crypto/rc4"
	"fmt"
)

func main() {
	key := []byte("rc4-secret-key")
	c, err := rc4.NewCipher(key)
	if err != nil {
		panic(err)
	}

	plaintext := []byte("hello RC4")
	ciphertext := make([]byte, len(plaintext))
	c.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("RC4 ciphertext: %x\n", ciphertext)
}
