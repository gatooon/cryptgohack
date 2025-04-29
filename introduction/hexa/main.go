package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main(){
	if len(os.Args) < 2 {
		log.Fatal("Argument is needed")
	}
	hexa := os.Args[1]

	fmt.Printf("[o] Hexa Input: '%v'\n", hexa)
	byte_array, err := hex.DecodeString(hexa)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[o] Converted to Byte array: '%v'\n", byte_array)

	fmt.Printf("[o] Byte array converted to string: '%v'\n", string(byte_array))
}
