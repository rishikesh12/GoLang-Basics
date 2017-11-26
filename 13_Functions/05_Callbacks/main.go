package main

import "fmt"

func main() {
	print([]int{2, 3, 4, 5, 6}, func(n int) {
		fmt.Println(n)

	})
}

func print(numbers []int, callbackfunction func(int)) {
	for _, v := range numbers {
		callbackfunction(v)
	}

}

//Callback : Passing a func as an argument
