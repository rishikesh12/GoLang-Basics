package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		// %q stands for utf-8
		fmt.Printf("%d \t %q \n", i, i)
	}
}
