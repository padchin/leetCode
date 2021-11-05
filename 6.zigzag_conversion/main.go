package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}

func convert(s string, numRows int) string {

	inLength := len(s) // допустимо только для ASCII
	inBytes := []byte(s)

	var mat [][]byte

	i := 0

	column := make([]byte, numRows)

start:
	for i < inLength {
		for x := 0; x < numRows; x++ {
			if i < inLength {
				column[x] = inBytes[i]
			} else {
				mat = append(mat, column)
				break start
			}

			i++
		}

		mat = append(mat, column)
		column = make([]byte, numRows)

		for x := numRows - 2; x > 0; x-- {
			if i < inLength {
				column[x] = inBytes[i]
			} else {
				mat = append(mat, column)
				break start
			}

			i++
		}

		mat = append(mat, column)
		column = make([]byte, numRows)
	}

	var res []byte

	matDepth := len(mat)

	for row := 0; row < numRows; row++ {
		for col := 0; col < matDepth; col++ {
			symbol := mat[col][row]
			if symbol != 0 {
				res = append(res, symbol)
			}
		}
	}

	return string(res)
}
