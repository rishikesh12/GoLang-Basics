package main

import "fmt"

func main() {
	var ar [10]int
	for i := 0; i < len(ar); i++ {
		ar[i] = i
	}
	for i := 0; i < len(ar); i++ {
		fmt.Println(ar[i])
	}
}

//There are no dynamic arrays in go
//arrays have [size] whereas slices are empty[]
