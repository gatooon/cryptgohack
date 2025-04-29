package main

import (
	"fmt"
)

func main(){

	byte_array := []byte{99, 114, 121, 112, 116, 103, 111, 104, 97, 99, 107}

	fmt.Printf("[o] Input: '%v'\n", byte_array)
	fmt.Printf("[o] Output converted to string: '%v'\n", string(byte_array))
}
