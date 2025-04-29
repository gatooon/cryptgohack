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

	fmt.Printf("[o] Input Big Int: '%v'\n", bigint_argument)

	hexa_converted := fmt.Sprintf("%x", bigint_argument)

	hexa_array, err := hex.DecodeString(hexa_converted)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[o] Big int converted to Hexa: '%v'\n", hexa_array)

	fmt.Printf("[o] Hexa array converted to string: '%v'\n", string(hexa_array))

}
