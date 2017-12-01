package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go increment("Foo:")
	go increment("Bar:")
	wg.Wait()
}

func increment(s string) {
	for i := 0; i < 20; i++ {

		atomic.AddInt64(&counter, 1)
		fmt.Println(s, counter)
	}
	wg.Done()
}

//Atomicity is another way to prevent race conditions(atomic package)
