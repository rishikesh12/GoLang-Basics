package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First string
	Last  string
}

func main() {
	p1 := Person{"Krishna", "Ravishankar"}
	json.NewEncoder(os.Stdout).Encode(p1) //NewEncoder returns a pointer of type Encoder
}
