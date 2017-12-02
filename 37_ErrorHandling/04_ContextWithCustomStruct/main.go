package main

import (
	"fmt"
)

type negativeerror struct {
	value float32
	err   error
}

func (e *negativeerror) Error() string { //negativeerror must implement Error() interface to create a custom error struct
	return fmt.Sprintf("%v %v", e.value, e.err)
}
func main() {
	_, err := squareroot(-10)
	if err != nil {
		fmt.Println(err)
	}
}

func squareroot(val float32) (float32, error) {
	if val < 0 {
		nge := fmt.Errorf("negative number")
		return 0, &negativeerror{val, nge}
	}
	return 0, nil
}

//see use of error structs in standard library net/dail.go, net/net.go, json/decode.go
