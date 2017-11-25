package main

import "fmt"

var b string = "Hello" //package scope

func main() {
	a := 10
	c := true
	d := 'w'
	//zero values
	var e string
	var f bool
	var g int
	var h float32
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%c \n", d)
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)

}
