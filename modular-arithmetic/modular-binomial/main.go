package main

import (
	"fmt"
	"log"
	"math/big"
	"errors"
)

func main(){
	N := ""

	a1 := ""
	b1 := ""
	e1 := ""
	c1 := ""

	a2 := "" 
	b2 := ""
	e2 := ""
	c2 := ""

	values_table, err := ParseValues(N, a1, b1, e1, c1, a2, b2, e2, c2)
	if err != nil {
		log.Fatal(err)
	}

	p := GetValue(values_table["a1"], values_table["a2"], values_table)
	q := GetValue(values_table["b1"], values_table["b2"], values_table)

	fmt.Printf("[o] q = %v\n", q.String())
	fmt.Printf("[o] p = %v\n", p.String())
}

func GetValue(x1 *big.Int, x2 *big.Int, values_table map[string]*big.Int) *big.Int{
	t1 := big.NewInt(0).Exp(x1, big.NewInt(0).Mul(big.NewInt(0).Neg(values_table["e1"]),values_table["e2"]), values_table["N"])
	t2 := big.NewInt(0).Exp(values_table["c1"], values_table["e2"], values_table["N"])

	t3 := big.NewInt(0).Exp(x2, big.NewInt(0).Mul(big.NewInt(0).Neg(values_table["e2"]),values_table["e1"]), values_table["N"])
	t4 := big.NewInt(0).Exp(values_table["c2"], values_table["e1"], values_table["N"])

	expr1 := big.NewInt(0).Mul(t1, t2)
	expr2 := big.NewInt(0).Mul(t3, t4)

	eliminated := big.NewInt(0).Sub(expr2, expr1)

	value := big.NewInt(0).GCD(nil, nil, eliminated, values_table["N"])

	return value

}

func ParseValues(N_string, a1_string, b1_string, e1_string, c1_string, a2_string, b2_string, e2_string, c2_string string) (map[string]*big.Int, error) {
	N, ok := big.NewInt(0).SetString(N_string, 10)
	if !ok {
		return nil, errors.New("Cant't convert to big int")
	}

	a1, ok := big.NewInt(0).SetString(a1_string, 10)
	if !ok {
		return nil, errors.New("Cant't convert to big int")
	}

	a2, ok := big.NewInt(0).SetString(a2_string, 10)
	if !ok {
		return nil, errors.New("Cant't convert to big int")
	}

	b1, ok := big.NewInt(0).SetString(b1_string, 10)
	if !ok {
		return nil, errors.New("Cant't convert to big int")
	}

	b2, ok := big.NewInt(0).SetString(b2_string, 10)
	if !ok {
		return nil, errors.New("Cant't convert to big int")
	}

	e1, ok := big.NewInt(0).SetString(e1_string, 10)
	if !ok {
		return nil, errors.New("Cant't convert to big int")
	}

	e2, ok := big.NewInt(0).SetString(e2_string, 10)
	if !ok {
		return nil, errors.New("Cant't convert to big int")
	}

	c1, ok := big.NewInt(0).SetString(c1_string, 10)
	if !ok {
		return nil, errors.New("Cant't convert to big int")
	}

	c2, ok := big.NewInt(0).SetString(c2_string, 10)
	if !ok {
		return nil, errors.New("Cant't convert to big int")
	}

	return map[string]*big.Int{
		"N": N,
		"a1": a1,
		"b1": b1,
		"e1": e1,
		"c1": c1,
		"a2": a2,
		"b2": b2,
		"e2": e2,
		"c2": c2,
	}, nil
}
