package main

import "fmt"

const s string = "Hello"
const s1 = "I am an untyped constant"
const x = 10
const (
	a = 12
	b = 23
	c = 34
)

const (
	d = iota //0
	e = iota //1
	f = iota //2
)
const (
	g = iota
	h
	i
)
const (
	j = iota
	k = 1 << (iota * 10) //Left shift
	l = 1 << (iota * 10)
)

func main() {
	const p = "Wassup"
	fmt.Println("s-", s)
	fmt.Println("p-", p)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Printf("%b\n", k)
	fmt.Printf("%b\n", l)
	fmt.Println(s1)
	y := x + 10
	fmt.Println(y)
}

//A constant is an unchanging value
//An iota is go's enumeration
//An untyped constant is not any type, hence it provides an advantage of being used in any operations without type conversion
