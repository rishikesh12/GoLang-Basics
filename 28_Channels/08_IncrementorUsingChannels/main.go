package main

import "fmt"

func main() {
	c1 := incrementor("Foo:", 10)
	c2 := incrementor("Bar:", 5)
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final sum of Foo", <-c3, "Final sum of Bar", <-c4)
}

func incrementor(s string, n int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for val := range c {
			sum += val
		}
		out <- sum
		close(out)
	}()
	return out

}
