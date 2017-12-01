package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	c := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(c)
	}()
	for val := range c {
		fmt.Println(val)
	}

}
