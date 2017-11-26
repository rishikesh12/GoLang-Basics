package main

import "fmt"

func main() {
	fmt.Println([]byte("Hello")) //convert to a array of byte
	fmt.Println([]byte(" 世界"))
	for i := 400; i < 1400; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}
}

//rune is a character
//rune is an alias for int32
//an interger value reperesenting utf value
