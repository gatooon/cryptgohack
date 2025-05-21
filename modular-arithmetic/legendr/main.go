package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

func main(){
	if len(os.Args) < 2 {
		log.Fatal("Argument is needed")
	}

 	file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
    }
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var p_string, numbers_string string
    for scanner.Scan() {
		if len(scanner.Text()) != 0 {
			if scanner.Text()[0] == []byte("p")[0]{
				p_string = strings.Split(scanner.Text(), " = ")[1]
			}
			if scanner.Text()[0] == []byte("i")[0]{
				numbers_string = strings.Split(scanner.Text(), " = ")[1]
			}
		}
    }

	numbers_string = numbers_string[1:]
	numbers_string = numbers_string[:len(numbers_string)-1]

	p := new(big.Int)
	p, ok := p.SetString(p_string, 10)
	if !ok {
		log.Fatal("Error Occured")
	}

	for _, number_string := range(strings.Split(numbers_string, ", ")){

		number := new(big.Int)
		number, ok := number.SetString(number_string, 10)
		if !ok {
			log.Fatal("Error Occured with big int conversion")
		}
		
		square_root, err := GetSquareRoot(number, p)
		if err != nil {
			fmt.Println(err)
			continue
		}

		square_root_2 := big.NewInt(0).Sub(p, square_root)
		fmt.Printf("[o] 2 Square root of %v mod %v is \n ------- \n %v\n ------- \n AND \n -------- \n %v \n -------- \n", number, p, square_root, square_root_2) 
	}

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

func GetSquareRoot(a *big.Int, p *big.Int) (*big.Int, error) {
	
	if CheckQResidue(a, p) != true {
		return big.NewInt(0), errors.New("[!] Can't apply algorithm, non quadratic residue")
	}

	fmt.Printf("[o] Quadratic residue\n")

	if big.NewInt(0).Mod(p, big.NewInt(4)).String() != "3" {
		return big.NewInt(0), errors.New("[!] Can't apply algorithm, p != 3 mod 4")
	}

	fmt.Printf("[o] Algorithm applicable\n")
	
	exponent := big.NewInt(0).Add(p, big.NewInt(1))
	exponent = big.NewInt(0).Div(exponent, big.NewInt(4))
	value := big.NewInt(0).Exp(a, exponent, p)

	return value, nil
}
