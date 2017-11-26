package main

import "fmt"

var x int

func main() {

	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

}
func wrapper() func() int {
	x = x + 1
	return func() int {
		return x
	}

}

//Interesting have to find out why it doesn't keep incrementing
