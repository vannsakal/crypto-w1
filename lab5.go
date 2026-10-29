package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	var text = "Hello"

	fmt.Println("Text:", text)

	fmt.Print("Binary: ")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%08b ", text[i])
	}
	fmt.Println()

	fmt.Println("Hex:", hex.EncodeToString([]byte(text)))
	fmt.Println("Base64:", base64.StdEncoding.EncodeToString([]byte(text)))
}
