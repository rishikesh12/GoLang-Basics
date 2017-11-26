package main

import "fmt"

func main() {
	func(n int) {
		fmt.Println(n)
	}(7)
}
