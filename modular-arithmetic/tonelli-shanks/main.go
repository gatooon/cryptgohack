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
	Q := big.NewInt(0).Sub(p, big.NewInt(1))
	S := big.NewInt(0)
	for big.NewInt(0).Mod(Q, big.NewInt(2)).String() == "0" {
		S = big.NewInt(0).Add(S, big.NewInt(1))
		Q = big.NewInt(0).Div(Q, big.NewInt(2))
		if Q.String() == "0" {
			log.Fatal("error")
		}
	}

	fmt.Printf("[o] odd prime for tonelli-shanks is %v\n", Q.String())
	fmt.Printf("[o] finding a non quadratic residu...\n")

	z := big.NewInt(2)
	for CheckQResidue(z, p) != false {
		z = big.NewInt(0).Add(z, big.NewInt(1))
	}
	fmt.Printf("[o] non quadratic residu for tonelli-shanks if %v\n", z.String())

	M := S
	c := big.NewInt(0).Exp(z, Q, p)
	t := big.NewInt(0).Exp(a, Q, p)
	R := big.NewInt(0).Exp(a, big.NewInt(0).Div(big.NewInt(0).Add(Q, big.NewInt(1)), big.NewInt(2)), p)

	for t.String() != "1" {
		i := big.NewInt(0)
		tmp := t
		for tmp.String() != "1" {
			i = big.NewInt(0).Add(i, big.NewInt(1))
			tmp = big.NewInt(0).Mod(big.NewInt(0).Mul(tmp, tmp), p)
		}

		pow2 := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(0).Sub(big.NewInt(0).Sub(M, i), big.NewInt(1)), nil)
		b := big.NewInt(0).Exp(c, pow2, p)

		M = i
		c = big.NewInt(0).Exp(b, big.NewInt(2), p)
		t = big.NewInt(0).Mod(big.NewInt(0).Mul(big.NewInt(0).Mul(t, b), b), p)
		R = big.NewInt(0).Mod(big.NewInt(0).Mul(R, b), p)
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

