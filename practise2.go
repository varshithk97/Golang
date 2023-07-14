package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {

	return a / b
}

func main() {
	var num1, num2 int
	var operator string

	fmt.Println("Simple Calculator")
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)
	fmt.Print("Enter the operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	var result int
	var err error

	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = subtract(num1, num2)
	case "*":
		result = multiply(num1, num2)
	case "/":
		result = divide(num1, num2)
	default:
		fmt.Println("Invalid operator")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

/* func main() {
var x, y = 10, 2

fmt.Printf("x + y = %d\n", x+y)
fmt.Printf("x - y = %d\n", x-y)
fmt.Printf("x * y = %d\n", x*y)
fmt.Printf("x / y = %d\n", x/y)
fmt.Printf("x mod y = %d\n", x%y)
var z = 5
fmt.Println(x < y && x > z)
fmt.Println(x < y || x > z)
fmt.Println(!(x == y && x > z)) */
