package main

import "fmt"

func main() {
	slice1 := []int{}        //shorthand way of making a slice
	var slice2 []int         //var way of making a slice
	slice3 := make([]int, 0) //make way of creating a slice
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice3 == nil)
	fmt.Println(slice2 == nil)
}

//Using make is the preferred method
