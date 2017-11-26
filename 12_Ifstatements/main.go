package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number")
	fmt.Scanf("%d", &n)
	if n%2 == 0 {
		fmt.Println("Entered value is even")
	} else {
		fmt.Println("Entered value is odd")
	}
}
