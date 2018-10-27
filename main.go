package main

import (
	"fmt"

	"github.com/liuyh73/DES-Golang/des"
)

func main() {
	var key string
	var clear_text string
	fmt.Print("please input secret key with 8 bits：")
	fmt.Scanf("%s\n", &key)
	fmt.Print("please input the clear text: ")
	fmt.Scanf("%s\n", &clear_text)
	cipher_text := des.Encrypt(clear_text, key)
	fmt.Println("明文加密后:", cipher_text)
	fmt.Println("密文解密后:", des.Decrypt(cipher_text, key))
}
