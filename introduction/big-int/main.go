package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"log"
	"math/big"
)

func main(){
	if len(os.Args) < 2 {
		log.Fatal("Argument is needed")
	}

	bigint_argument := new(big.Int)
	bigint_argument.SetString(os.Args[1], 10)

	fmt.Println(bigint_argument)

	hexa_converted := fmt.Sprintf("%x", bigint_argument)

	hexa_array, err := hex.DecodeString(hexa_converted)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(hexa_array))

}
