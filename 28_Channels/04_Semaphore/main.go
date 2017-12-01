package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	sema := make(chan bool)
	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}
		sema <- true
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}
		sema <- true
	}()
	go func() {
		<-sema
		<-sema
		close(c)
	}()
	/* wrong Way
	<-sema this will be waiting until a go routine is done but it causes a dealock because the value isn't consumed from the Channel
	*/
	for val := range c {
		fmt.Println(val)
	}
}
