package main

import "fmt"

func main() {
	fmt.Printf("%d %#x %#o \n", 42, 42, 42)
	fmt.Printf("%d %x %o \n", 42, 42, 42)
	fmt.Printf("%d %#X %#o", 42, 42, 42)
}
