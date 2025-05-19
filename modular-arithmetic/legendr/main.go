package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
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

	for _, number_string := range(strings.Split(numbers, ", ")){

		number_string

		fmt.Printf("[o] p = %v ; a = %v\n", p, number)
		
	}
    

	//	err, square_root := GetSquareRoot(a, p)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Printf("[o] Square Root of %v mod %v is %v", a, p, square_root)
	//
}

func CheckQResidue(a BigInteger, p BigInteger) bool{
	res := int(math.Pow(a, (p-1)/2))%int(p)
	if res == 1 {
		return true
	}
	return false
}

func GetSquareRoot(a float64, p float64) (error,int) {
	
	if CheckQResidue(a, p) != true {
		return errors.New("Can't apply algorithm, non quadratic residue"), 0
	}


	if int(p)%4 != 3 {
		return errors.New("Can't apply algorithm, p != 3 mod 4"), 0
	}

	return nil, int(math.Pow(a, (p+1)/4))%int(p)
}
