package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string `json:"-"`            //tag indicating to ignore this type
	Age   int    `json:"wisdom score"` //tag to change the display name
}

func main() {
	p1 := person{"Krishna", "Ravishankar", 21}
	bs, _ := json.Marshal(p1) //Marshal converts to json
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}

// Marshaling and Unmarshaling is used when working on a string or a slice of bytes
