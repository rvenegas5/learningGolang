package main

import "fmt"

func greet(name, lastname string) {
	fmt.Printf("Hello, %s %s from GO\n", name, lastname)
}

func isItEven(number int) string {
	if number%2 == 0 {
		return "The number is even"
	} else {
		return "the number is odd"
	}
}

func main() {
	var name string
	var lastname string
	fmt.Print("Write your name: ")
	fmt.Scan(&name)
	fmt.Print("Write your lastname: ")
	fmt.Scan(&lastname)
	greet(name, lastname)
	var number int
	fmt.Println("-------------- Second part --------------")
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	result := isItEven(number)
	fmt.Println(result)
}
