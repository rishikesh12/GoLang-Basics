package main

import "fmt"

func main() {
	avg := average(10, 20, 30, 40, 50)
	fmt.Println(avg)
	avg1 := []int{45, 23, 78, 90} //Passing a slice (ex for variadic arguments)
	fmt.Println(average(avg1...)) //Opening the slice
}
func average(si ...int) float32 {
	fmt.Println(si)
	fmt.Printf("%T\n", si)
	total := 0
	for _, v := range si {
		total += v
	}
	return float32(total) / float32(len(si))
}

//Variadic functions are functions that support variable number of arguments
