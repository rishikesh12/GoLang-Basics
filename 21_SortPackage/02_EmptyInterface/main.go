package main

import "fmt"

type vehicles interface {
}
type vehicle struct {
	seat  int
	color string
}
type car struct {
	vehicle
	speed int
}
type boat struct {
	vehicle
	steeringratio int
}

func main() {
	i10 := car{}
	to := boat{}
	ti := vehicle{}
	in := []vehicles{i10, to, ti}
	fmt.Printf("%T\n", in)
	for _, value := range in {
		fmt.Println(value)
	}
}

//Empty Interface accepts any type
