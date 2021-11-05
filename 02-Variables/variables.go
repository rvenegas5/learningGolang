package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Ways to declare variables in golang

	//First way
	var name string = "William"

	// Second way
	var lastname = "Venegas"

	// Cool way
	age := 22

	fmt.Println("Hello, My name is " + name + " " + lastname + " and I'm " + strconv.Itoa(age) + " years old.")

	// Declare variables but these have initial values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// For constants variables we have the following
	const pi = 3.141592

	fmt.Println(pi)

}
