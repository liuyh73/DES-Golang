package des

func iteration(clear_text_init_replace string, keys []string) string {
	L := make([]string, 0)
	R := make([]string, 0)
	L = append(L, clear_text_init_replace[:32])
	R = append(R, clear_text_init_replace[32:])
	for k := 0; k < 16; k++ {
		// Extended replacement
		R_extended := extendedReplacement(R[k])
		// xor with keys[k]
		R_extended_xor := xorWithKeys_K(R_extended, keys[k])
		// S-box transfer
		R_extended_xor_S_trans := sBoxTransfer(R_extended_xor)
		// P-box transfer
		R_extended_xor_S_P_trans := pBoxTransfer(R_extended_xor_S_trans)
		// xor with L[k]
		R_extended_xor_S_P_trans_xor := xorWithL_K(R_extended_xor_S_P_trans, L[k])

		L = append(L, R[k])
		R = append(R, R_extended_xor_S_P_trans_xor)
	}
	R_16_L_16 := R[16] + L[16]
	return R_16_L_16
}

func extendedReplacement(R_K string) string {
	extended_replacement_matrix := getExtendedReplacementMatrix()
	width := len(extended_replacement_matrix)
	height := len(extended_replacement_matrix[0])
	R_extended := ""
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			R_extended += string(R_K[extended_replacement_matrix[i][j]-1])
		}
	}
	return R_extended
}

func xorWithKeys_K(R_extended, keys_K string) string {
	R_extended_xor := ""
	for i := 0; i < len(keys_K); i++ {
		R_extended_xor += string((byte(R_extended[i]-'0') ^ byte(keys_K[i]-'0')) + '0')
	}
	return R_extended_xor
}

func sBoxTransfer(R_extended_xor string) string {
	R_extended_xor_S_trans := ""
	for i := 0; i < 8; i++ {
		R_extended_xor_slice := R_extended_xor[6*i : 6*(i+1)]
		row := getRow(R_extended_xor_slice[0], R_extended_xor_slice[5])
		column := getColumn(R_extended_xor_slice[1:5])
		S_trans_data := getSBoxN(i)[row][column]
		// fmt.Printf("Row: %d, Column: %d ", row, column)
		// fmt.Printf("%d \n", S_trans_data)
		R_extended_xor_S_trans += decimalToBinary(S_trans_data)
	}
	return R_extended_xor_S_trans
}

func pBoxTransfer(R_extended_xor_S_trans string) string {
	R_extended_xor_S_P_trans := ""
	p_box := getPBox()
	for i := 0; i < len(p_box); i++ {
		R_extended_xor_S_P_trans += string(R_extended_xor_S_trans[p_box[i]-1])
	}
	return R_extended_xor_S_P_trans
}

func xorWithL_K(R_extended_xor_S_P_trans, L_K string) string {
	R_extended_xor_S_P_trans_xor := ""
	for i := 0; i < len(L_K); i++ {
		R_extended_xor_S_P_trans_xor += string((byte(R_extended_xor_S_P_trans[i]-'0') ^ byte(L_K[i]-'0')) + '0')
	}
	return R_extended_xor_S_P_trans_xor
}
