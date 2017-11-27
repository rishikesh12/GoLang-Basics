package main

import "fmt"

func main() {
	m := make(map[string]int) //creates a map of string key type and int value type
	m["k1"] = 10
	m["k2"] = 15
	fmt.Println(m)
	v := m["k1"]
	fmt.Println(v)
	fmt.Println("Length:", len(m))
	delete(m, "k2") //Deleting an entry
	fmt.Println(m)
	_, ok := m["k2"] //ok is initialized to true if the key is present in the map
	fmt.Println(ok)

	n := map[string]int{"foo": 1, "bar": 2} //alternate way to create a map without using make function
	fmt.Println(n)
}
