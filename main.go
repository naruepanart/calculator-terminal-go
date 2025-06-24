package main

import "fmt"

func main() {
	var a, b int
	var choice int

	fmt.Println("Version 1.1")
	
	for {
		fmt.Println("Choose your operation:")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("0. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice < 0 || choice > 3 {
			fmt.Println("Invalid choice. Please choose a number between 0 and 3.")
			continue
		}

		if choice == 0 {
			fmt.Println("Bye! Have a great day!")
			break
		}

		fmt.Print("Enter the first number: ")
		fmt.Scan(&a)
		fmt.Print("Enter the second number: ")
		fmt.Scan(&b)

		if choice == 1 {
			fmt.Println("Result:", sum(a, b))
		}
		if choice == 2 {
			fmt.Println("Result:", subtract(a, b))
		}
		if choice == 3 {
			fmt.Println("Result:", multiply(a, b))
		}
	}
}

func sum(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}