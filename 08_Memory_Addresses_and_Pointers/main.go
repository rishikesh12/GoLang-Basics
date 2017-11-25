package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Value of x:", x)
	fmt.Println("Memory Address of x:", &x)
	fmt.Printf("%d\n", &x)

	var y *int = &x
	fmt.Println(y)
	fmt.Println(*y)
	fmt.Printf("%T\n", y)
	*y = 42
	fmt.Println(*y)
	increment(x)
	increment1(&x)
	fmt.Println(*y)
}

func increment(x int) {
	x = x + 1
	fmt.Printf("%p\n", &x) //different memory address
}

func increment1(x *int) {
	*x = *x + 1
}

//*int is its own type
