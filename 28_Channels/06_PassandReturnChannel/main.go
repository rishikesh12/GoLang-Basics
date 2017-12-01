package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)
	for val := range cSum {
		fmt.Println(val)
	}
}

func incrementor() chan int {
	out := make(chan int)
	go func() {

		for i := 1; i <= 10; i++ {
			out <- i //Putting values into the channel
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int

		for val := range c { //Pulling of the channel
			sum += val
		}
		out <- sum
		close(out)

	}()
	return out
}
