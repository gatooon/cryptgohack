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
	
	hex_sentence := os.Args[1]


	bytes_sentence, err := hex.DecodeString(hex_sentence)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 256; i++ {
		fmt.Printf("[o] Byte %v(%b) %v\n",i, i, string(xor(bytes_sentence, byte(i))))
	}
}

func xor(sentence []byte, key byte) []byte{
	output := []byte{}
	for _, sentence_element := range sentence {
		xored := sentence_element^key
		output = append(output, xored)
	}
	return output
}
