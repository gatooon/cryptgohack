package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
)

func main(){
	output_table := []string{} //output file as array

	p := "" // P used in the encryption algorithm
	tmp_string := ""
	ascii_table := []string{}
	i := 1
	for _, integer := range output_table {
		value, err := CheckQuadratic(integer, p)
		if err != nil {
			log.Fatal(err)
		}
		if value {
			tmp_string = tmp_string + "1"
		}else{
			tmp_string = tmp_string + "0"
		}
		if i == 8 {
			i = 1 
			ascii_table = append(ascii_table, tmp_string)
			tmp_string = ""
		}else{
			i++
		}
	}

	flag, err := BinaryToTxt(ascii_table)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("[o] Flag: %v\n", flag)
}

func CheckQuadratic(a string, p string) (bool, error){
	big_a, ok := big.NewInt(0).SetString(a, 10)
	if !ok {
		return false, errors.New("Can't convert string to bigint") 
	}
	big_p, ok := big.NewInt(0).SetString(p, 10)
	if !ok {
		return false, errors.New("Can't convert string to bigint") 
	}

	value := big.NewInt(0).Exp(big_a, big.NewInt(0).Div(big.NewInt(0).Sub(big_p, big.NewInt(1)),big.NewInt(2)), big_p)
	if value.String() == "1" {
		return true, nil
	}else{
		return false, nil
	}
}

func BinaryToTxt(binary_array []string) (string, error){
	decimal_array := []byte{}
	for _, binary := range binary_array {
		value, err := strconv.ParseInt(binary, 2, 64)
		if err != nil {
			return "", err
		}
		decimal_array = append(decimal_array, byte(value))
	}

	return string(decimal_array), nil
}












