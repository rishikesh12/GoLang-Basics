package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := gen()
	c2 := factorial(c1)
	c3 := factorial(c1)
	c4 := merge(c3, c2)
	for val := range c4 {
		fmt.Println(val)
	}
}

func gen() chan int {
	out := make(chan int)
	go func() {
		for j := 1; j <= 100; j++ {
			for i := 1; i <= 10; i++ {
				out <- i
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
func merge(c ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(c))
	for _, val := range c {
		go func(n chan int) {
			for val1 := range n {
				out <- val1
			}
			wg.Done()
		}(val)

	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
