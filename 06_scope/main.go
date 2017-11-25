package main

import "fmt" //file level scope

var a int = 42

func main() {
	fmt.Println(a)
	foo()
}

func foo() {
	fmt.Println("From Foo", a) //example of package level scope
}
