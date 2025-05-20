package main

import (
//	"bufio"
	"errors"
	"fmt"
	"log"
//	"math"
	"math/big"
	"os"
//	"strings"
)

func main(){
	if len(os.Args) < 3 {
		log.Fatal("Argument is needed")
	}

	p := new(big.Int)
	p, ok := p.SetString(os.Args[1], 10)
	if !ok {
		log.Fatal("Error Occured")
	}

	a := new(big.Int)
	a, ok = p.SetString(os.Args[2], 10)
	if !ok {
		log.Fatal("Error Occured")
	}


	value, err := GetSquareRoot(a, p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)



// 	file, err := os.Open(os.Args[1])
//    if err != nil {
//        fmt.Println(err)
//    }
//	defer file.Close()
//	scanner := bufio.NewScanner(file)
//	var p_string, numbers_string string
//    for scanner.Scan() {
//		if len(scanner.Text()) != 0 {
//			if scanner.Text()[0] == []byte("p")[0]{
//				p_string = strings.Split(scanner.Text(), " = ")[1]
//			}
//			if scanner.Text()[0] == []byte("i")[0]{
//				numbers_string = strings.Split(scanner.Text(), " = ")[1]
//			}
//		}
//    }
//
//	numbers_string = numbers_string[1:]
//	numbers_string = numbers_string[:len(numbers_string)-1]
//
//	p := new(big.Int)
//	p, ok := p.SetString(p_string, 10)
//	if !ok {
//		log.Fatal("Error Occured")
//	}
//
//	for _, number_string := range(strings.Split(numbers, ", ")){
//
//		number_string
//
//		fmt.Printf("[o] p = %v ; a = %v\n", p, number)
//		
//	}
    

	//	err, square_root := GetSquareRoot(a, p)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Printf("[o] Square Root of %v mod %v is %v", a, p, square_root)
	//
}

func CheckQResidue(a *big.Int, p *big.Int) bool{
	//res := int(math.Pow(a, (p-1)/2))%int(p)

	res := big.NewInt(0).Sub(p, big.NewInt(1))
	fmt.Println(res.String())
	res = big.NewInt(0).Div(res, big.NewInt(2))
	fmt.Println(res.String())
	res = big.NewInt(0).Exp(a, res, nil)


	if res.String() == "1" {
		return true
	}
	return false
}

func GetSquareRoot(a *big.Int, p *big.Int) (*big.Int, error) {
	
	if CheckQResidue(a, p) != true {
		return big.NewInt(0), errors.New("Can't apply algorithm, non quadratic residue")
	}


	if big.NewInt(0).Mod(p, big.NewInt(4)).String() != "3" {
		return big.NewInt(0), errors.New("Can't apply algorithm, p != 3 mod 4")
	}

	
	exponent := big.NewInt(0).Add(p, big.NewInt(1))
	exponent = big.NewInt(0).Div(exponent, big.NewInt(4))
	value := big.NewInt(0).Exp(a, exponent, nil)
	value = big.NewInt(0).Mod(value, p)

	return value, nil
}
