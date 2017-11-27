package main

import "fmt"

func main() {
	sl := make([]int, 0, 20)
	for i := 0; i < 70; i++ {
		sl = append(sl, i)
		fmt.Println("len:", len(sl), " capacity:", cap(sl), " value:", sl[i])
	}
	s2 := make([]int, 3, 40)
	s2[0] = 7
	s2[1] = 9
	s2[2] = 10
	s2 = append(s2, sl...)
	fmt.Println(s2)
}
