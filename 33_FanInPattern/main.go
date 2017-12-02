package main

import "fmt"

//A simple fan-in example program
//fan-in: takes input from multiple channels and writes it onto a single channel(multiplexing)
func main() {
	c1 := write("Foo:")
	c2 := write("Bar:")
	fan := fanin(c1, c2)
	for val := range fan {
		fmt.Println(val)
	}
}

func write(s string) chan string {
	out := make(chan string)
	go func() {
		out <- fmt.Sprint(s, 1)
		close(out)
	}()

	return out
}

func fanin(c1, c2 chan string) chan string {
	out := make(chan string)
	done := make(chan bool)
	go func() {
		for val := range c1 {
			out <- val

		}
		done <- true

	}()
	go func() {
		for val := range c2 {
			out <- val
		}
		done <- true
	}()
	go func() {
		<-done
		<-done
		close(out)
	}()
	return out
}
