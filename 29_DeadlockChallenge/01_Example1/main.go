package main

import "fmt"

func main() {
	c := make(chan int)

	//c <- 1 causes deadlock since the control stops here and there is nothing to recieve the value from the channel
	go func() { //think of go routines as they are separate entities running somewhere
		c <- 1
	}()
	fmt.Println(<-c)
}
