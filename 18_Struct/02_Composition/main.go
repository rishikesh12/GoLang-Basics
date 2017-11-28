package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p person) fullname() string { //p person is a reciver and any value of type person will have access to fullname method
	return p.first + " " + p.last
}
func main() {
	var p1 person
	p1 = person{"Krishna", "Ravishankar"}
	fmt.Println(p1.fullname())
}
