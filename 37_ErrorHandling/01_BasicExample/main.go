package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	lf, err := os.Create("37_ErrorHandling/01_BasicExample/log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(lf)
}
func main() {
	_, err := os.Open("notextfile.txt")
	if err != nil {
		//fmt.Println(err)
		log.Println(err)
		//log.Fatalln(err)
		//panic(err)
	}
}

//Println just prints the error to the standard output
// log.Println is similar to Println but it also provides date and time of the error
//panic gives more information regarding trace
// Fatalln returns a non zero exit status on error
