package main

import (
	"fmt"
	"log"
	"encoding/hex"
)

func main() {

	s_box_revers := [256]string{}
	matrix := [4][4]byte{}

	matrix_reversed, err := SubBytes(matrix, s_box_revers)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(Matrix2Byte(matrix_reversed))
}

func Byte2Matrix(sequence [16]byte) [4][4]byte {
var returned_array [4][4]byte
	var converted [4]byte

	n := 0
	for i := 0; i < 4; i++ {
		copy(converted[:], sequence[n:n+4])
		returned_array[i] = converted
		n = n + 4
	}

	return returned_array
}

func Matrix2Byte(matrix [4][4]byte) [16]byte {
	return [16]byte{matrix[0][0], matrix[0][1], matrix[0][2], matrix[0][3],
		matrix[1][0], matrix[1][1], matrix[1][2], matrix[1][3],
		matrix[2][0], matrix[2][1], matrix[2][2], matrix[2][3],
		matrix[3][0], matrix[3][1], matrix[3][2], matrix[3][3]} //c'est pas beau mais c'est plus perfs
}

func CreateRoundedKeys(message [4][4]byte, key [4][4]byte) [11][4][4]byte {
	var returned_matrix [11][4][4]byte
	for i := 0; i < 11; i++ {
		if i == 0 {
			returned_matrix[i] = Xor2Matrix(message, key)
		} else {
			returned_matrix[i] = Xor2Matrix(returned_matrix[i-1], key)
		}
	}
	return returned_matrix
}

func Xor2Matrix(m1 [4][4]byte, m2 [4][4]byte) [4][4]byte {
	var returned_matrix [4][4]byte

	for x, line := range m1 {
		for y, unit := range line {
			returned_matrix[x][y] = unit ^ m2[x][y]
		}
	}

	return returned_matrix
}

func SubBytes(matrix [4][4]byte, s_box [256]string) ([4][4]byte, error) {
	var returned_matrix [4][4]byte
	for x, line := range(matrix){
		for y, value := range(line){
			output_hex := s_box[value]
      output_byte, err := hex.DecodeString(output_hex)
    	if err != nil {
    	    return returned_matrix, err
    	}
			returned_matrix[x][y] = output_byte[0]
		}
	}

	return returned_matrix, nil

}

