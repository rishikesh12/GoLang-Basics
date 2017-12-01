package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go foo("foo:")
	go bar("bar:")
	wg.Wait()
}

func foo(s string) {
	for i := 1; i < 20; i++ {
		mutex.Lock()
		counter++
		fmt.Println(s, counter)
		time.Sleep(time.Duration(5 * time.Millisecond))
		mutex.Unlock()
	}
	wg.Done()
}

func bar(s string) {
	for i := 1; i < 30; i++ {
		mutex.Lock()
		counter++
		fmt.Println(s, counter)
		time.Sleep(time.Duration(20 * time.Millisecond))
		mutex.Unlock()
	}
	wg.Done()
}

//Multiple processes accessing and modifying the same variable
