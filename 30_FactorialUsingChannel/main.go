package main

import "fmt"

func main() {
	fmt.Println("Enter the Value of N")
	var n int
	//n = 4
	fmt.Scanf("%d", &n)
	c1 := incrementor(n)
	fact := factorial(c1)
	for val := range fact {
		fmt.Println(val)
	}
}

func incrementor(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(c chan int) chan int {
	out := make(chan int)
	go func() {
		f := 1
		for val := range c {
			f *= val
		}
		out <- f
		close(out)
	}()
	return out

}
