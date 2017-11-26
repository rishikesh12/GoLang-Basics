package main

import "fmt"

func main() {
	filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		if n > 1 {
			return true
		} else {
			return false
		}
	})
}
func filter(numbers []int, callbackfunction func(int) bool) {
	for _, v := range numbers {
		if callbackfunction(v) {
			fmt.Println(v)
		}
	}
}
