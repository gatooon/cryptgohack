package main

import (
	"fmt"
)

func main(){
	slice := [][]int{{2, 5}, {3, 11}, {5, 17}}
	fmt.Printf("[o]%v\n", ChineseRemainder(slice))
}

func ExtendedEuclidian(a int, b int) int {
	input_a, input_b := a, b
	var mod_result, t, q int
	t1 := 0
	t2 := 1
	for {
//		fmt.Printf("[o][o] a = %v, b = %v, q = %v, mod = %v, t1 = %v, t2 = %v, t = %v\n", a, b, q, mod_result, t1, t2, t)
		if b == 0 {
			t1 = modNegative(t1, input_a)
			fmt.Printf("[o] Modular Inverse for %v mod %v is %v\n", input_b, input_a, t1)
			return t1
		}
		q = a/b
		t = t1 - t2 * q
		t1 = t2
		t2 = t

		mod_result = a%b
		a = b
		b = mod_result
	}
}

func ChineseRemainder(given [][]int) int {
	M := 1
	var M1, M_1, output int
	for _, element := range given {
		M = M*element[1]
	}

	fmt.Printf("[o]M = %v\n", M)

	for _, element := range given {
		M1 = M/element[1]
		M_1 = ExtendedEuclidian(element[1], M1)
		output = output + (element[0] * M1 * M_1)
		fmt.Printf("[o]a = %v, Mx = %v, Modular Inverse = %v, output = %v\n", element[0], M1, M_1, output)
	}

	return output%M
}

func modNegative(a int, b int) int{
	return (a % b + b) % b
}
