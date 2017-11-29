package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person
	bs := []byte(`{"First":"Krishna","Last":"Ravishankar","Age":21}`)
	json.Unmarshal(bs, &p1)
	fmt.Println(p1.Age)
}
