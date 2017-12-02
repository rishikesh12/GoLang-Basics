package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := gen(2, 3)
	//Two workers to generate square
	c2 := square(c1) //fan out
	c3 := square(c1) //fan out
	//Merge the results
	c4 := merge(c2, c3) //fan in

	for val := range c4 {
		fmt.Println(val)
	}

}

func gen(n ...int) chan int {
	fmt.Printf("%T\n", n)
	out := make(chan int)
	go func() {
		for _, val := range n {
			out <- val
		}
		close(out)
	}()
	return out
}

func square(c chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(ch ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ch))
	fmt.Printf("%T\n", ch)
	for _, val := range ch {
		go func(c chan int) {
			for n := range c {
				out <- n
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

//A very nice example to illustrate the power of concurrency
// We have 2 workers which perform the task of squaring
