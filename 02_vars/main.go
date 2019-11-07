package main

import "fmt"

func main() {
	// Using var
	// var name string = "Bernard"
	var age int32 = 27
	var isCool = true

	isCool = false

	// Using const
	const isSeriouslyCool = true

	// shorthand for var
	anothername := "BJ"
	size := 1.3

	// Multiple assignments
	name, email := "Bernard", "letslearntechyt@gmail.com"

	fmt.Println(name, age, isCool, isSeriouslyCool)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Println(anothername, size, email)
}
