package des

import (
	"encoding/hex"
	"strconv"
)

func decimalToBinary(data int) string {
	switch data {
	case 0:
		return "0000"
	case 1:
		return "0001"
	case 2:
		return "0010"
	case 3:
		return "0011"
	case 4:
		return "0100"
	case 5:
		return "0101"
	case 6:
		return "0110"
	case 7:
		return "0111"
	case 8:
		return "1000"
	case 9:
		return "1001"
	case 10:
		return "1010"
	case 11:
		return "1011"
	case 12:
		return "1100"
	case 13:
		return "1101"
	case 14:
		return "1110"
	case 15:
		return "1111"
	}
	return ""
}

func hexadecimalToBinary(data byte) string {
	switch data {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'a':
		return "1010"
	case 'b':
		return "1011"
	case 'c':
		return "1100"
	case 'd':
		return "1101"
	case 'e':
		return "1110"
	case 'f':
		return "1111"
	}
	return ""
}

func reverse(keys []string) []string {
	keys_reverse := make([]string, 0)
	for i := len(keys) - 1; i >= 0; i-- {
		keys_reverse = append(keys_reverse, keys[i])
	}
	return keys_reverse
}

func formatKey(key string) string {
	hexKey := hex.EncodeToString([]byte(key))
	binKey := ""
	for i := 0; i < len(hexKey); i++ {
		binKey += hexadecimalToBinary(hexKey[i])
	}
	return binKey
}

func binText(text string) string {
	binText := ""
	for i := 0; i < len(text); i++ {
		binText += hexadecimalToBinary(text[i])
	}
	return binText
}

func hexText(text string) string {
	hexText := ""
	for i := 0; i < len(text)/4; i++ {
		dec_text, _ := strconv.ParseInt(text[i*4:i*4+4], 2, 64)
		hexText += strconv.FormatInt(dec_text, 16)
	}
	return hexText
}
