package main

import "fmt"

func main() {
	sliceofslice := make([][]string, 0, 10)
	student1 := make([]string, 2)
	student1[0] = "Ravishankar"
	student1[1] = "Krishna"
	sliceofslice = append(sliceofslice, student1)
	student2 := make([]string, 2)
	student2[0] = "Jani"
	student2[1] = "Harsh"
	sliceofslice = append(sliceofslice, student2)
	fmt.Println(sliceofslice)

}
