package main

import "fmt"

func main() {
	var n int
	n = 10
	c := make(chan int)
	sema := make(chan bool)
	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
	for i := 1; i <= n; i++ {
		go func() {
			for val := range c {
				fmt.Println(val)
			}
			sema <- true
		}()
	}
	for i := 1; i <= n; i++ {
		<-sema
	}
}
