package main

import (
	"fmt"

	"github.com/liuyh73/DES-Golang/des"
)

func main() {
	key := "0001001100110100010101110111100110011011101111001101111111110001"
	clear_text := "0010010110100011010001011100011001101001101010111100110111101111"
	cipher_text := des.Encrypt(clear_text, key)
	fmt.Println(cipher_text)
	plain_text := des.Decrypt(cipher_text, key)
	fmt.Println(plain_text)
	fmt.Println(clear_text)
}
