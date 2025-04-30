package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main(){
	if len(os.Args) < 5 {
		log.Fatal("Argument is needed")
	}
	
	hex_key1 := os.Args[1]
	hex_sentence1 := os.Args[2]
	hex_sentence2 := os.Args[3]
	hex_sentence3 := os.Args[4]


	bytes_key1, err := hex.DecodeString(hex_key1)
	if err != nil {
		log.Fatal(err)
	}
	bytes_sentence1, err := hex.DecodeString(hex_sentence1)
	if err != nil {
		log.Fatal(err)
	}
	bytes_sentence2, err := hex.DecodeString(hex_sentence2)
	if err != nil {
		log.Fatal(err)
	}
	bytes_sentence3, err := hex.DecodeString(hex_sentence3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[o] key 1: %x\n", bytes_key1)
	bytes_key2 := xor(bytes_sentence1, bytes_key1)
	fmt.Printf("[o] key 2: %x\n", bytes_key2)
	bytes_key3 := xor(bytes_sentence2, bytes_key2)
	fmt.Printf("[o] key 3: %x\n", bytes_key3)

	key3_xored := xor(bytes_sentence3, bytes_key3)
	fmt.Printf("[o] sentence 3 xored with key3: %x\n", key3_xored)

	key2_xored := xor(key3_xored, bytes_key2)
	fmt.Printf("[o] sentence 3 (xored with key3) xored with key 2: %x\n", key2_xored)

	key1_xored := xor(key2_xored, bytes_key1)
	fmt.Printf("[o] sentence 3 (xored with key3 and key 2) xored with key 1: %x\n", key1_xored)

	fmt.Printf("[o] Flag: %v\n", string(key1_xored))

}

func xor(sentence []byte, key []byte) []byte{
	output := []byte{}
	for i, sentence_element := range sentence {
		xored := sentence_element^key[i]
		output = append(output, xored)

	}
	return output
} 
