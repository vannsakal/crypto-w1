package main

import (
	"fmt"
	"os"
	"os/exec"
)

func run(file string) {
	cmd := exec.Command("go", "run", file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	var choice int
	for {
		fmt.Println("\nChoose lab to run (1-6, 0 to exit):")
		fmt.Scan(&choice)
		if choice == 0 {
			break
		}
		switch choice {
		case 1:
			run("lab1.go")
		case 2:
			run("lab2.go")
		case 3:
			run("lab3.go")
		case 4:
			run("lab4.go")
		case 5:
			run("lab5.go")
		case 6:
			run("lab6.go")
		default:
			fmt.Println("Invalid choice")
		}
	}
}
