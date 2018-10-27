package des

func initialReplace(clear_text string) string {
	clear_text_init_replace := ""
	initial_replace_matrix := getInitialReplaceMatrix()
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			clear_text_init_replace += string(clear_text[initial_replace_matrix[i][j]-1])
		}
	}
	return clear_text_init_replace
}
