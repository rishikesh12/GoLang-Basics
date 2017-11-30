package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}
type calculate interface {
	area() float32
}
type calculate1 interface {
	perimeter() float32
}

func (c circle) area() float32 { //Value Reciver
	return float32(math.Pi * c.radius * c.radius)
}

func (c *circle) perimeter() float32 { //Pointer Reviver
	return float32(2 * math.Pi * c.radius)
}
func printarea(c calculate) {
	fmt.Println(c.area())

}
func printperimeter(c calculate1) {
	fmt.Println(c.perimeter())

}
func main() {
	c := circle{10.00}
	printarea(c)
	printperimeter(&c)
	printarea(&c)
}

//Recivers are of two types value type and the pointer types
//with the value type you can pass the value and the address
//with the pointer type you have to pass the address
