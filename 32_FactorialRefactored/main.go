package main

import "fmt"

func main() {
	c1 := gen()
	f := factorial(c1)
	for val := range f {
		fmt.Println(val)
	}
}

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			for j := 1; j <= 10; j++ {
				out <- j
			}
		}
		close(out)

	}()
	return out
}

func factorial(c chan int) chan int {
	out := make(chan int)
	go func() {
		for val := range c {
			out <- fact(val)
		}

		close(out)

	}()
	return out
}

func fact(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	return f
}
