package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"encoding/base64"
)

func main(){
	if len(os.Args) < 2 {
		log.Fatal("Argument is needed")
	}
	hexa := os.Args[1]

	fmt.Printf("[o] Hexa input: '%v'\n", hexa)
	byte_array, err := hex.DecodeString(hexa)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[o] Hexa converted to byte array: '%v'\n", byte_array)

	b64_string := base64.StdEncoding.EncodeToString(byte_array)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[o] Byte array converted to base64: '%v'\n", b64_string)
}
