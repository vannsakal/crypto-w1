package main

import "fmt"

func main() {
	var a = 10
	var b = 5

	fmt.Println("both positive:", a > 0 && b > 0)
	fmt.Println("one greater than the other:", a > b || b > a)
	fmt.Println("not equal:", !(a == b))
}
