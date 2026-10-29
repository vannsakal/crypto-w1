package main

import "fmt"

func myAND(a int, b int) {
	fmt.Printf("AND: %b\n", a&b)
}

func myOR(a int, b int) {
	fmt.Printf("OR: %b\n", a|b)
}

func myXor(a int, b int) {
	fmt.Printf("XOR: %b\n", a^b)
}

func myNOT(a int, b int) {
	fmt.Printf("NOT a: %b\n", ^a)
	fmt.Printf("NOT b: %b\n", ^b)
}

func myLeftShift(a int) {
	fmt.Printf("LEFT SHIFT: %b\n", a<<2)
}

func myRightShift(a int) {
	fmt.Printf("RIGHT SHIFT: %b\n", a>>2)
}

func main() {
	var a = 0b1001
	var b = 0b1100

	myXor(a, b)
	myNOT(a, b)
	myOR(a, b)
	myAND(a, b)
	myLeftShift(a)
	myRightShift(a)

	var c = a
	c &= b
	fmt.Println("c &= b:", c)
	c |= b
	fmt.Println("c |= b:", c)
	c ^= b
	fmt.Println("c ^= b:", c)
	c <<= 2
	fmt.Println("c <<= 2:", c)
	c >>= 2
	fmt.Println("c >>= 2:", c)
}
