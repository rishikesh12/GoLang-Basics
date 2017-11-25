package main

import "fmt"

func main() {
	x := 10 //outer scope encloses the inner scope
	fmt.Println(x)
	{
		y := 5
		fmt.Println(x)
		fmt.Println(y)
	}
}
