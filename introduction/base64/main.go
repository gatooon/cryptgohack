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

	byte_array, err := hex.DecodeString(hexa)
	if err != nil {
		log.Fatal(err)
	}

	b64_string := base64.StdEncoding.EncodeToString(byte_array)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b64_string)
}
