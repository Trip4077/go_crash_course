package main

import "fmt"

func main() {
	// map[key type] value type
	emails := make(map[string]string)

	//Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete
	delete(emails, "Bob")

	fmt.Println(emails)

	// Assign at creation
	users := map[int]string{0: "Bob", 1: "Sharon"}

	fmt.Println(users)
	fmt.Println(users[0])
}
