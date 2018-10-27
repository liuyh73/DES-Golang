package des

func reverseReplace(R_16_L_16 string) string {
	// reverse replace
	cipher_text := ""
	reverse_replace := getReverseReplace()
	for i := 0; i < len(reverse_replace); i++ {
		cipher_text += string(R_16_L_16[reverse_replace[i]-1])
	}
	return cipher_text
}
