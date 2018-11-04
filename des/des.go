package des

import (
	"encoding/hex"
	"fmt"
	"os"
)

func Encrypt(clear_text, key string) string {
	extra := 8 - len(clear_text)%8
	for i := 0; i < extra; i++ {
		clear_text = clear_text + string('0'+extra)
	}
	clear_text = hex.EncodeToString([]byte(clear_text))
	return hexText(des(clear_text, key, true))
}

func Decrypt(cipher_text, key string) string {
	clear_text_hex := hexText(des(cipher_text, key, false))
	clear_text, _ := hex.DecodeString(clear_text_hex)
	clear_text_len := len(clear_text)
	return string(clear_text[:clear_text_len-int(clear_text[clear_text_len-1]-'0')])
}

func des(text, key string, tag bool) string {
	if len(key) != 8 {
		fmt.Println("The secret key need to be 8 bits.")
		os.Exit(0)
	}
	key = formatKey(key)
	keys := getKeys(key)
	final_text := ""
	if !tag {
		keys = reverse(keys)
	}
	for i := 0; i < len(text)/16; i++ {
		textSub := binText(text[i*16 : i*16+16])
		text_init_replace := initialReplace(textSub)
		R_16_L_16 := iteration(text_init_replace, keys)
		final_text += reverseReplace(R_16_L_16)
	}
	return final_text
}
