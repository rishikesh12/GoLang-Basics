package main

import "fmt"

func main() {
	var m map[string]string // do not create a map like this as it returns a nil instance
	//m["k1"] = "Harsh"       //Can't assign and it throws an error
	//m["k2"] = "Jani"        //No append function is available for map
	fmt.Println(m)
	var n = map[string]string{} //composite literal declaration, it works
	n["k1"] = "Krishna"
	n["k2"] = "Chethan"
	fmt.Println(n)

	for key, val := range n {
		fmt.Println("key-", key, " Value-", val)
	}
}

//Always use make or composite literal declaration
