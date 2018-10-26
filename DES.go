package main

import (
	"fmt"
)

// 10101010 11110000 11101011 11110111 00010101 00110011 10100101 01101010
const clear_text = "1010101011110000111010111111011100010101001100111010010101101010"
const initial_matrix_side = 8

var PC_2_matrix = [7][8]int{
	{49, 41, 33, 25, 17, 9, 1},
	{58, 50, 42, 34, 26, 18, 10},
	{2, 59, 51, 43, 35, 27, 19},
	{11, 3, 60, 52, 44, 36, 57},
	{31, 23, 15, 7, 62, 54, 46, 38},
	{30, 22, 14, 6, 61, 53, 45, 37},
	{29, 21, 13, 5, 28, 20, 12, 4},
}

func main() {
	// for i := 0; i < initial_matrix_side; i++ {
	// 	for j := 0; j < initial_matrix_side; j++ {
	// 		fmt.Printf("%d ", initial_matrix[i][j])
	// 	}
	// 	fmt.Println()
	// }

	// initial replacement
	clear_text_init_replace := ""
	for i := 0; i < initial_matrix_side; i++ {
		for j := 0; j < initial_matrix_side; j++ {
			clear_text_init_replace += string(clear_text[initial_matrix[i][j]-1])
		}
	}
	L0 := clear_text_init_replace[:32]
	L1 := clear_text_init_replace[32:]

	// Encryption processing
	// PC-1(子密钥的生成)

}
