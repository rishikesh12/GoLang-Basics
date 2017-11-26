package main

import "fmt"

func main() {
	m := make([]string, 1, 24) //type,length,capacity
	fmt.Println(m)
	change(m)
	fmt.Println(m)
}
func change(n []string) {
	n[0] = "Hello"
}

//make is used to make a slice which is a refernce type
