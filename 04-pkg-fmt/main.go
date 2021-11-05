package main

import "fmt"

func main() {
	hello := "Hello"
	world := "World"

	fmt.Print(hello)
	fmt.Println(" ", world)

	name := "William"
	lastName := "Venegas"
	age := 22

	// This is for print a string with format
	fmt.Printf("Hello my name is %s %s and I'm %d years old\n", name, lastName, age)
	// If I don't know what is the format of my varibles
	fmt.Println("------------ If I don't know the datatype ------------")
	fmt.Printf("Hello my name is %v %v and I'm %v years old\n", name, lastName, age)

	// If I wanna save data in a varible
	fmt.Println("------------ If I save the message in a variable ------------")
	message := fmt.Sprintf("Hello my name is %s %s and I'm %d years old\n", name, lastName, age)
	fmt.Println(message)

	// If I wanna print the datatype of my variable
	fmt.Printf("The datatype for message variable is: %T\n", message)

	// Use it for data input
	fmt.Println("------------ input ------------")
	var data string
	fmt.Print("Write something: ")
	fmt.Scanln(&data)
	fmt.Println("You wrote this => ", data)
}
