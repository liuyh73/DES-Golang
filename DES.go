package main

import (
	"fmt"

	"github.com/liuyh73/DES-Golang/sipher_code"
)

// 10101010 11110000 11101011 11110111 00010101 00110011 10100101 01101010
const clear_text = "1010101011110000111010111111011100010101001100111010010101101010"
const initial_siphercode_side = 8

var sipherCode *sipher_code.Sipher_code

func main() {
	clear_text_init_replace := ""
	// initial replacement
	initial_replace_matrix := sipherCode.GetInitialReplaceMatrix()
	for i := 0; i < initial_siphercode_side; i++ {
		for j := 0; j < initial_siphercode_side; j++ {
			clear_text_init_replace += string(clear_text[initial_replace_matrix[i][j]-1])
		}
	}
	// fmt.Println(clear_text_init_replace)
	// Encryption processing
	// PC-1(子密钥的生成)
	for i := 1; i <= 16; i++ {
		sipherCodeN := sipherCode.GetSipherCodeN(i)
		fmt.Println(sipherCodeN)
	}
}
