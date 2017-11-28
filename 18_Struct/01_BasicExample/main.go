package main

import "fmt"

type person struct { //defining a person struct type
	first string
	last  string
	age   int
}

func main() {
	var p1 person
	//p1 := person{}
	p1.first = "Krishna"
	p1.last = "Ravishankar"
	p1.age = 21
	p2 := person{"Harsh", "Jani", 21} //alternate way to define value(struct literal)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
