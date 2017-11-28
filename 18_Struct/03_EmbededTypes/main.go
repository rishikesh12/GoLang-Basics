package main

import "fmt"

type person struct {
	first string
	last  string
}
type average struct {
	person //Embedding person type in average
	avg    int
}

func main() {
	avg1 := average{
		person: person{first: "Krishna", last: "Ravishankar"}, //Clear way to declare values
		avg:    10,
	}
	fmt.Println(avg1.avg)
	fmt.Println(avg1.person.first + " " + avg1.person.last)
	fmt.Println(avg1.first) //inner value is promoted to be accesible by the outer type
	fmt.Println(avg1)
	fmt.Printf("%T\n", avg1)
}
