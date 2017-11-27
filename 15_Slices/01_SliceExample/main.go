package main

import "fmt"

func main() {
	var sl = []int{1, 2, 3, 4, 5} //slice
	fmt.Println(sl)
	fmt.Println(sl[2:4]) //slicing a slice
	fmt.Println(sl[3])   //index access
}
