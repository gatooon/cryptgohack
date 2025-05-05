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

	number1_args, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	number2_args, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	a := number1_args
	b := number2_args

	euclidian_table := GenEuclidianTable(a,b)

	for _, line := range(euclidian_table) {
		fmt.Printf("[o] %v = %v*%v+%v\n",line["a"], line["q"], line["b"],line["r"])
	}
}

func EclDivision(numerator int, denominator int) (int, int) {
	return numerator/denominator, numerator%denominator
}

func GenEuclidianTable(a int, b int) []map[string]int {
	euclidian_table := []map[string]int{}

	for {
		quotien, remainder := EclDivision(a, b)
		euclidian_table = append(euclidian_table, map[string]int{"a": a, "b": b, "q": quotien, "r": remainder})
		if remainder == 0 {
			euclidian_table = euclidian_table[:len(euclidian_table) - 1]
			break
		}
		a = b
		b = remainder
	}

	return euclidian_table
}

func ReverseEuclidianTable(euclidian_table []map[string]int) (int, int){
	table_len := len(euclidian_table)
	for i := table_len; i >= 0; i-- {
		line := euclidian_table[i]
		
	}
}
