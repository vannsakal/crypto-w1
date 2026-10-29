package main

import (
	"encoding/hex"
	"fmt"
)

func xorEncrypt(text string, key byte) string {
	var result = make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key
	}
	return string(result)
}

func main() {
	var text = "Hello"
	var key byte = 'K'

	ciphertext := xorEncrypt(text, key)
	decrypted := xorEncrypt(ciphertext, key)

	fmt.Println("Original:", text)
	fmt.Println("Ciphertext (hex):", hex.EncodeToString([]byte(ciphertext)))
	fmt.Println("Decrypted:", decrypted)
}
