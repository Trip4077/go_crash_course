package main

import "fmt"

func main() {
	ids := []int{123, 321, 231, 132}

	// Loop through IDs
	for i, id := range ids {
		fmt.Printf("%d - ID %d\n", i, id)
	}

	// Loop through IDs, no index
	for _, id := range ids {
		fmt.Printf("ID %d\n", id)
	}

	// Add IDs together
	sum := 0

	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum: ", sum)

	// Range With Map
	users := map[int]string{0: "Bob", 1: "Sharon", 2: "mike"}

	for key, user := range users {
		fmt.Printf("%d: %s\n", key, user)
	}

	for key := range users {
		fmt.Printf("ID: %d\n", key)
	}

	for _, value := range users {
		fmt.Println("User: " + value)
	}
}
