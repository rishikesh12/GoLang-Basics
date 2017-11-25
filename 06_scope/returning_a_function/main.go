package main

import "fmt"

func main() {
	value := increment()
	fmt.Println(value())
	fmt.Println(value())
}

func increment() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}

// We can return a function type by making use of anonymous function
