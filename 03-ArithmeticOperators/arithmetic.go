package main

import "fmt"

func main() {

	a := 20
	b := 10

	// Sum
	result := a + b
	fmt.Println("Sum = ", result)

	// Subtraction
	result = a - b
	fmt.Println("Subtraction = ", result)

	// Multiplication
	result = a * b
	fmt.Println("Multiplication = ", result)

	// Division
	result = a / b
	fmt.Println("Division = ", result)

	// Decimal division
	var anotherOne float64 = 3.0 / 2.0
	fmt.Println("Decimal division = ", anotherOne)

	// Module
	result = a % b
	fmt.Println("Module", result)
}
