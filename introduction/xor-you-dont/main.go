package main

import (
	"encoding/hex"
	"log"
	"os"
	"fmt"
)

func main(){
	if len(os.Args) < 4 {
		log.Fatal("Argument is needed")
	}

	script_type := os.Args[1]

	sentence := os.Args[2]
	known_word := os.Args[3]

	switch script_type {
	case "0":
		TrySentence(sentence, known_word)
	case "1":
		TryDecode(sentence, known_word)
	default:
		log.Fatal("First argument have to be 0 (try to find key) or 1 (try with key)")
	}

}

func SingleByteXor(sentence []byte, key byte) []byte{
	output := []byte{}
	for _, sentence_element := range sentence {
		xored := sentence_element^key
		output = append(output, xored)
	}
	return output
}

func Xor(sentence []byte, key []byte) []byte {
	key_len := len(key)
	i := 0
	output := []byte{}
	for _, sentence_element := range sentence {
		if i == key_len{
			i = 0
		}
		xored := sentence_element^key[i]
		output = append(output, xored)
		i ++
	}
	return output
}

func TrySentence(sentence string, known_word string) {
	bytes_sentence, err := hex.DecodeString(sentence)
	if err != nil {
		log.Fatal(err)
	}

	
	first_ascii := 65
	ascii_number := 60

	fmt.Printf("[o] List of ascii letter that will be used (ascii_number %v): \n", ascii_number)
	for i := 0; i < ascii_number; i++ {
		fmt.Printf("%v", string(first_ascii+i))
	}
	fmt.Printf("\n")

	for j, letter := range known_word{
		for i := 0; i < ascii_number; i++ {
			ascii_letter := byte(first_ascii+i)
			xored_sentence := SingleByteXor(bytes_sentence, ascii_letter)
			if string(xored_sentence[j]) == string(letter) {
				fmt.Printf("[o] For '%v' position(%v), it may be '%v'\n", j, string(xored_sentence[j]), string(ascii_letter))
				break
			}
		}
	}

}

func TryDecode(sentence, key string) {
	bytes_sentence, err := hex.DecodeString(sentence)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[o] Try with key: %v\n", key)

	output := Xor(bytes_sentence, []byte(key))

	fmt.Printf("[o] Output: %v\n", string(output))

}
