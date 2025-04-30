package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main(){
	if len(os.Args) < 3 {
		log.Fatal("Argument is needed")
	}

	sentence := os.Args[1]
	number := os.Args[2]

	int_number, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}

	bytes_sentence := []byte(sentence)
	bytes_number := byte(int_number)

	output := []byte{}
	for _, element := range bytes_sentence {
		xored := bytes_number^element
		fmt.Printf("[o] char: %v(%b), number: %v(%b), xored: %v(%b)\n",element, element,bytes_number, bytes_number, xored, xored)

		output = append(output, xored)
	}
	fmt.Printf("[o] Output string: %v\n", string(output))

}
