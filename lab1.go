package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var a = 20
	var b = 10

	a = b
	fmt.Println("a =", a)

	a += b
	fmt.Println("a += b:", a)

	a -= b
	fmt.Println("a -= b:", a)

	a *= b
	fmt.Println("a *= b:", a)

	a /= b
	fmt.Println("a /= b:", a)

	a %= b
	fmt.Println("a %= b:", a)
}
