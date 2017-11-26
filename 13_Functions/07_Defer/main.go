package main

import "fmt"

func main() {
	defer world()
	hello()
}

func world() {
	fmt.Println("World")
}
func hello() {
	fmt.Println("hello")
}

//defer executes the function before exiting main
//defer can be used to close opened files before exit
