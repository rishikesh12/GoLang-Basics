package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int { //functions from Interface interface in sort package that the type must implement
	return len(p)
}

func (p people) Less(i, j int) bool {
	if p[i] < p[j] {
		return true
	}
	return false
}

func (p people) Swap(i, j int) {
	temp := p[i]
	p[i] = p[j]
	p[j] = temp
}

func main() {
	studygroup := people{"Harsh", "Chethan", "Krishna", "Kaushik", "Charan"}
	for i := 0; i < studygroup.Len(); i++ {
		for j := i + 1; j < studygroup.Len(); j++ {
			if !studygroup.Less(i, j) {
				studygroup.Swap(i, j)
			}
		}
	}
	studygroup2 := people{"Ritesh", "Giridhar", "Gautham", "Eshwar"}
	sort.Sort(studygroup2)
	fmt.Println(studygroup2)
	fmt.Println(studygroup)

	//Interfaces are defined only for type but what happe ns if it is a variable
	names := []string{"Bibek", "Kanthi", "Arnab"}

	sort.StringSlice(names).Sort() //Converts slice of strings to stringslice then sort method can be requested
	fmt.Println(names)
	//sort.Sort(sort.StringSlice(names))
	sort.Sort(sort.Reverse(sort.StringSlice(names)))
	fmt.Println(names)
}
