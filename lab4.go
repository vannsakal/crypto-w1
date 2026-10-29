package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int) {
	if b == 0 {
		fmt.Println("Error: division by zero")
		return
	}
	fmt.Println("Result:", a/b)
}

func mod(a int, b int) {
	if b == 0 {
		fmt.Println("Error: division by zero")
		return
	}
	fmt.Println("Result:", a%b)
}

func main() {
	var choice int
	var a, b int

	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("1) Add 2) Sub 3) Mul 4) Div 5) Mod 6) Exit")
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		if choice == 6 {
			fmt.Println("Exited.")
			break
		}

		fmt.Print("Enter a: ")
		fmt.Scan(&a)

		fmt.Print("Enter b: ")
		fmt.Scan(&b)

		switch choice {
		case 1:
			fmt.Println("Addition:", add(a, b))
		case 2:
			fmt.Println("Subtraction:", sub(a, b))
		case 3:
			fmt.Println("Product:", mul(a, b))
		case 4:
			div(a, b)
		case 5:
			mod(a, b)
		default:
			fmt.Println("Invalid choice")
		}
	}
}
