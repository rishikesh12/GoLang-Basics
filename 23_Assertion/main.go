package main

import "fmt"

func main() {
	var a interface{} = "Wayan"
	str, ok := a.(string)
	if ok {
		fmt.Println(str)
	} else {
		fmt.Println("Not Ok")
	}
	var b interface{} = 7
	fmt.Println(b.(int) + 3)
}

//Assertion is only used for interfaces to assert certain types
