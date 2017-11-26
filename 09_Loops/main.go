package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	j := 0

	for j < 10 { // while like loop
		fmt.Println(j)
		j = j + 1

	}

	k := 0
	for {
		fmt.Println(k)
		k = k + 1
		if k >= 10 {
			break
		}

	}
	//nested loop
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, " - ", j)
		}
	}
}

//go has only for loop
