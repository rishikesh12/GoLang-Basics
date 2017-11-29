package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First string
	Last  string
	age   int //not exported
}

func main() {
	var p1 Person
	rdr := strings.NewReader(`{"First":"Krishna","Last":"Ravishankar"}`)
	json.NewDecoder(rdr).Decode(&p1) //NewDecoder requires a reader parameter and returns a pointer of Decoder type
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.age)
}
