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

	byte_array, err := hex.DecodeString(hexa)
	if err != nil {
		log.Fatal()
	}

	fmt.Println(string(byte_array))
}
