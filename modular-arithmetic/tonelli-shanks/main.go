package main

import (
	"fmt"
	"log"
	"os"
	"math/big"
)


func main(){
	if len(os.Args) < 3 {
		log.Fatal("Argument is needed")
	}

	p, ok := big.NewInt(0).SetString(os.Args[1], 10)
	if !ok {
		log.Fatal("Error occured during string conversion to big int")
	}

	a, ok := big.NewInt(0).SetString(os.Args[2], 10)
	if !ok {
		log.Fatal("Error occured during string conversion to big int")
	}

	if !CheckQResidue(a, p) {
		log.Fatal("Arguments are not quadratic residu")
	}

	square_root := TonelliShanks(a, p)

	fmt.Printf("[o] Square root of %v mod %v is %v\n", a.String(), p.String(), square_root.String())
}




func TonelliShanks(a *big.Int, p *big.Int) *big.Int{
	big_int := big.NewInt(0)
	Q := big_int.Sub(p, big.NewInt(1))
	fmt.Println(Q)
	S := big.NewInt(0)
	fmt.Println(S)
	for big_int.Mod(Q, big.NewInt(2)).String() == "0" {
		S = big.NewInt(0).Add(S, big.NewInt(1))
		Q = big.NewInt(0).Div(Q, big.NewInt(2))
		fmt.Println(Q)
		fmt.Println(S)
		if Q.String() == "0" {
			log.Fatal("error")
		}
	}

	fmt.Printf("[o] odd prime for tonelli-shanks is %v\n", Q.String())
	fmt.Printf("[o] finding a non quadratic residu...\n")

	z := big.NewInt(2)
	for CheckQResidue(z, p) != false {
		z = big_int.Add(z, big.NewInt(1))
	}
	fmt.Printf("[o] non quadratic residu for tonelli-shanks if %v\n", z.String())

	M := S
	c := big_int.Exp(z, Q, p)
	t := big_int.Exp(a, Q, p)
	R := big_int.Exp(a, big_int.Div(big_int.Add(Q, big.NewInt(1)), big.NewInt(2)), p)

	for t.String() != "1" {
		i := big.NewInt(0)
		tmp := t
		for tmp.String() != "1" {
			i = big_int.Add(i, big.NewInt(1))
			tmp = big_int.Mod(big_int.Mul(tmp, tmp), p)
		}

		pow2 := big_int.Exp(big.NewInt(2), big_int.Sub(big_int.Sub(M, i), big.NewInt(1)), nil)
		b := big_int.Exp(c, pow2, p)

		M = i
		c = big_int.Exp(b, big.NewInt(2), p)
		t = big_int.Mod(big_int.Mul(big_int.Mul(t, b), b), p)
		R = big_int.Mod(big_int.Mul(R, b), p)
	}

	return R
}


func CheckQResidue(a *big.Int, p *big.Int) bool{

	res := big.NewInt(0).Sub(p, big.NewInt(1))
	res = big.NewInt(0).Div(res, big.NewInt(2))
	res = big.NewInt(0).Exp(a, res, p)

	if res.String() == "1" {
		return true
	}
	return false
}

