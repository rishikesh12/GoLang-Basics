package main

import "fmt"

func main() {
	/*c := nums(2, 3)
	out := square(c)
	fmt.Println(<-out)
	fmt.Println(<-out)*/
	for n := range square(square(nums(2, 3))) { //pipeline to consume output
		fmt.Println(n)
	}
}

func nums(n ...int) chan int {
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
		for val := range c {
			out <- val * val
		}
		close(out)
	}()
	return out
}

/*A pipeline is a series of stages connected by channels, where each stage is a group of goroutines running the same function. In each stage, the goroutines

receive values from upstream via inbound channels
perform some function on that data, usually producing new values
send values downstream via outbound channels*/
