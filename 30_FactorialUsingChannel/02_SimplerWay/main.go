package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the value of n")
	fmt.Scanf("%d", &n)
	fact := factorial(n)
	for val := range fact {
		fmt.Println(val)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		f := 1
		for i := 1; i <= n; i++ {
			f *= i
		}
		out <- f
		close(out)
	}()
	return out
}
