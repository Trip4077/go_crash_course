package main

import "fmt"

func main() {
	// Long Method
	i := 1

	for i <= 10 {
		fmt.Println(i)
		// i = i + 1
		i++
	}

	// Short Method
	for j := 1; j <= 10; j++ {
		fmt.Printf("Number: %d\n", j)
	}

	// Fizzbuzz
	for k := 1; k <= 100; k++ {
		if k%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if k%5 == 0 {
			fmt.Println("Buzz")
		} else if k%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(k)
		}
	}
}
