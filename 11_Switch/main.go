package main

import "fmt"

func main() {
	var x string
	fmt.Println("Choose between fallon or kimmel")
	fmt.Scanf("%s", &x)
	switch x {
	case "fallon":
		fmt.Println("You have choosen Fallon")
		fallthrough
	case "kimmel":
		fmt.Println("You have choosen Kimmel")
	default:
		fmt.Println("You have not choosen anyone")

	}
	switchontype(7)
}
func switchontype(x interface{}) {
	switch x.(type) { //Asserting x is os type
	case int:
		fmt.Println("You have given an int")
	}
}

//No default fallthrough in go
