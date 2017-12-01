package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) //Unbuffered Channel
	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(time.Second)
}

//The channel is blocked until the value put int the channel is consumed by some other process
