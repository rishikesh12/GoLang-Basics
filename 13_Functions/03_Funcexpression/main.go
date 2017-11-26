package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("Hello")
	}
	greeting()
	fmt.Printf("%T\n", greeting)

	compliment := func() func() string {
		return func() string {
			return "Great"
		}
	}
	s := compliment()
	fmt.Println(s())
	fmt.Printf("%T\n", s)
}
