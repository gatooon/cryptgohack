package main

import (
	"fmt"
)

func main(){
	fmt.Printf("[o]%v\n", ExtendedEuclidian(13, 11))
}

func ExtendedEuclidian(a int, b int) int {
	fmt.Printf("[o]Extended Euclidian Algorithm\n")
	var mod_result, t, q int
	t1 := 0
	t2 := 1
	for {
		fmt.Printf("[o][o] a = %v, b = %v, q = %v, mod = %v, t1 = %v, t2 = %v, t = %v\n", a, b, q, mod_result, t1, t2, t)
		if b == 0 {
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
