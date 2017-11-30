package main

import (
	"fmt"
	"math"
)

type square struct {
	side int
}

type circle struct {
	radius float32
}

type shape interface {
	area() float32
}

func (s square) area() float32 { //square type implementing the shape interface
	return float32(s.side * s.side)
}
func (c circle) area() float32 {
	return float32(math.Pi * c.radius * c.radius)
}

func info(s shape) { //function that has a parameter of shape interface
	//fmt.Println(s)
	fmt.Println(s.area())
}
func main() {
	s := square{10}
	info(s) //you can pass the type that implements the interface
	c := circle{10}
	info(c)
}

//Circle and Square are concrete types and the interface shape shows no concrete behaviour
//Interface is an abstract type
