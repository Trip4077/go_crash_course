package main

import "fmt"

func main() {
	// arrays
	// fixed length and types
	var fruitArray [2]string

	// Assign values
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	fmt.Println(fruitArray)
	fmt.Println(fruitArray[0], fruitArray[1])

	// declare and assign
	fruits := [2]string{"Apple", "Orange"}

	fmt.Println(fruits)
	fmt.Println(fruits[0], fruits[1])

	//Slices
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
	fmt.Println(fruitSlice[1:3])
}
