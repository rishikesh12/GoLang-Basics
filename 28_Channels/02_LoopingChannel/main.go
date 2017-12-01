package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}
		close(c) //close signifies that nothing else can be added into the channel
	}()
	for val := range c { //but content from the channel can be consumed using range
		fmt.Println(val)
	}
}

// the for loop will wait for a value to be put in the channel and then it will consume it and then the control goes back to the writer
