package main

import "fmt"

func main() {
	foo := 'a'
	fmt.Println(foo)
	fmt.Println(string(foo))
	fmt.Println(rune(foo))
	fo := "a"
	fmt.Println(fo)
	fmt.Printf("%T\n", fo)

	foam := `This is the
	end hold your breath and
	count to ten` //raw string literal
	fmt.Println(foam)

}

//string is just a collection of runes
