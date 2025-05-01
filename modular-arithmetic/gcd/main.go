package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
)

func main(){
	if len(os.Args) < 3 {
		log.Fatal("Argument is needed")
	}

	number1_args, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	number2_args, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	number1 := number1_args
	number2 := number2_args
	var gcd int
	for {
		quotien, remainder := EclDivision(number1, number2)
		fmt.Printf("[o] %v / %v = %v reminder %v\n", number1, number2, quotien, remainder)
		if remainder == 0 {
			break
		}
		number1 = number2
		number2 = remainder
		gcd = remainder
	}

	fmt.Printf("[o] GCD of %v and %v is %v\n",number1_args, number2_args, gcd)

}

func EclDivision(numerator int, denominator int) (int, int) {
	return numerator/denominator, numerator%denominator
}
