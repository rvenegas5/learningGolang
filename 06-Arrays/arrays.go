package main

import "fmt"

func main() {
	// Declare an array, values are initialized to zero
	var numbers [5]int
	fmt.Println(numbers)

	// Adding values to the array
	numbers[0] = 5
	numbers[1] = 10
	numbers[2] = 20

	fmt.Println(numbers)
	// Print a specific position
	fmt.Println(numbers[4])

	// Initializing the array
	names := [3]string{"William", "Richard", "Melisa"}
	fmt.Println(names)

	// Initializing an array with unknown length
	colors := [...]string{
		"red",
		"blue",
		"green",
		"gold",
	}

	fmt.Println(colors, "with length ", len(colors))

	// Array with index
	coins := [...]string{
		0: "Dolar",
		2: "Euro",
		5: "Libra",
	}

	fmt.Println(coins, "with length of ", len(coins))
	fmt.Println(coins[0])
	coins[1] = "Peso"
	fmt.Println("----------- Using for -----------")
	for i := 0; i < len(coins); i++ {
		fmt.Println(coins[i])
	}

}
