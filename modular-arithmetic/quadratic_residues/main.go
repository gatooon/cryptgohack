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

	p, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	check_number, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < p; i++ {
		 res := i*i%p
		 fmt.Printf("[o] square(%v)(%v) mod %v == %v\n", i, i*i, p, res)
		 if res == check_number {
			 fmt.Printf("[o] %v is quadratic residues of %v mod %v\n", i, check_number, p)
		 }
	}
}
