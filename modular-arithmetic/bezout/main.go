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

	fmt.Printf("-----Reverse-----\n")

	reverse_euclidian_table := ReverseEuclidianTable(euclidian_table)
	for _, line := range(reverse_euclidian_table) {
		fmt.Printf("[o] = (%v)*%v+(%v)*%v\n",line["u"], line["a"], line["v"],line["b"])
	}

	u := reverse_euclidian_table[len(reverse_euclidian_table)-1]["u"]
	v := reverse_euclidian_table[len(reverse_euclidian_table)-1]["v"]
	gcd := euclidian_table[len(euclidian_table)-1]["r"]

	fmt.Printf("[o] u*%v+u*%v=GCD(%v,%v) \n[o] %v*%v+%v+%v=%v\n",number1_args, number2_args, number1_args, number2_args, u, number1_args, v, number2_args,gcd)
	fmt.Printf("[o] u=%v v=%v\n",u, v)

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

func ReverseEuclidianTable(euclidian_table []map[string]int) []map[string]int{
	table_len := len(euclidian_table)-1
	reverse_euclidian_table := []map[string]int{}
	var u, v int
	for i := table_len; i > 0; i-- {
		line := euclidian_table[i]
		subline := euclidian_table[i-1]

		if table_len == i {
			u = -(line["q"])
			v = u*-(subline["q"])+1
		}else{
			u = reverse_euclidian_table[len(reverse_euclidian_table)-1]["v"]
			v = u*-(subline["q"])+reverse_euclidian_table[len(reverse_euclidian_table)-1]["u"]
		}

		reverse_euclidian_table = append(reverse_euclidian_table, map[string]int{"a":subline["a"],"b":subline["b"], "u":u, "v":v})
	}

	return reverse_euclidian_table
}

