package main

import "fmt"

func main() {
	fmt.Println(getname("Kevin", "Zayn"))
}
func getname(p, q string) (string, string) {
	return fmt.Sprint(p, " ", q), fmt.Sprint(q, " ", p) //Sprint prints to a string rather than to the standard output
}
