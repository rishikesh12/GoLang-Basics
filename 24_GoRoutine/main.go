package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 1; i < 20; i++ {
		fmt.Println("foo:", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 1; i < 30; i++ {
		fmt.Println("bar:", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}

// go keyword is used to run processes concurrently
