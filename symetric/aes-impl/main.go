package main

import (
	"errors"
)

func main(){
}


func Byte2Matrix(sequence []byte) ([][]byte, error) {
	returned_array := [][]byte{}

	if len(sequence) != 16 {
		return nil, errors.New("Can't convert a non 16 byte sequence")
	}

	n := 0
	for i := 0; i < 4; i++ {
		returned_array = append(returned_array, sequence[n:n+4])
		n = n+4
	}

	return returned_array, nil
}

func Matrix2Byte(matrix [][]byte) ([]byte, error) {
	n := 0
	for i, values := range(matrix){
		if len(values) != 4{
			return nil, errors.New("Can't convert a non 4x4 matrix to bytes sequence")
		}
		n = i
	}
	if n != 3 {
		return nil, errors.New("Can't convert a non 4x4 matrix to bytes sequence")
	}
	
	return []byte{matrix[0][0], matrix[0][1], matrix[0][2], matrix[0][3], 
				matrix[1][0], matrix[1][1], matrix[1][2], matrix[1][3],
				matrix[2][0], matrix[2][1], matrix[2][2], matrix[2][3],
				matrix[3][0], matrix[3][1], matrix[3][2], matrix[3][3]}, nil //c'est pas beau mais c'est plus perfs
}
