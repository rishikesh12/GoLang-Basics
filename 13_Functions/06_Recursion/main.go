package main

import "fmt"

func main() {
	var n int = 7
	fact := factorial(n)
	fmt.Println(fact)
}
func factorial(n int) int {

	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}

}

//Recursion : A function which calls itself
