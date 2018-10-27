package des

func Encrypt(clear_text, key string) string {
	return des(clear_text, key, true)
}

func Decrypt(cipher_text, key string) string {
	return des(cipher_text, key, false)
}

func des(text, key string, tag bool) string {
	keys := getKeys(key)

	if !tag {
		keys = reverse(keys)
	}
	text_init_replace := initialReplace(text)
	R_16_L_16 := iteration(text_init_replace, keys)
	final_text := reverseReplace(R_16_L_16)
	return final_text
}
