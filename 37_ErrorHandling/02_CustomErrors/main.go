package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float32) (float32, error) {
	if f < 0 {
		return 0, errors.New("Number is less than zero")
	}
	return f * f, nil
}

//Error values in go are just like any other values
