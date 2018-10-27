package sipher_code

var initial_sipher_code = "0001001100110100010101110111100110011011101111001101111111110001"

var initial_replace_matrix = [8][8]int{
	{58, 50, 42, 34, 26, 18, 10, 2},
	{60, 52, 44, 36, 28, 20, 12, 4},
	{62, 54, 46, 38, 30, 22, 14, 6},
	{64, 56, 48, 40, 32, 24, 16, 8},
	{57, 49, 41, 33, 25, 17, 9, 1},
	{59, 51, 43, 35, 27, 19, 11, 3},
	{61, 53, 45, 37, 29, 21, 13, 5},
	{63, 55, 47, 39, 31, 23, 15, 7},
}

var pc_1 = [56]int{
	57, 49, 41, 33, 25, 17, 9,
	1, 58, 50, 42, 34, 26, 18,
	10, 2, 59, 51, 43, 35, 27,
	19, 11, 3, 60, 52, 44, 36,
	63, 55, 47, 39, 31, 23, 15,
	7, 62, 54, 46, 38, 30, 22,
	14, 6, 61, 53, 45, 37, 29,
	21, 13, 5, 28, 20, 12, 4,
}

var pc_2 = [48]int{
	14, 17, 11, 24, 1, 5,
	3, 28, 15, 6, 21, 10,
	23, 19, 12, 4, 26, 8,
	16, 7, 27, 20, 13, 2,
	41, 52, 31, 37, 47, 55,
	30, 40, 51, 45, 33, 48,
	44, 49, 39, 56, 34, 53,
	46, 42, 50, 36, 29, 32,
}

var displacements = [16]int{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

type Sipher_code struct {
}

func (*Sipher_code) GetInitialReplaceMatrix() [8][8]int {
	return initial_replace_matrix
}

func (*Sipher_code) getCN(round int) []int {
	displacement := getDisplacements(round)
	CN := make([]int, 28-displacement)
	copy(CN, pc_1[displacement:28])
	for i := 0; i < displacement; i++ {
		CN = append(CN, pc_1[i])
	}
	return CN
}

func (*Sipher_code) getDN(round int) []int {
	displacement := getDisplacements(round)
	DN := make([]int, 28-displacement)
	copy(DN, pc_1[28+displacement:])
	for i := 0; i < displacement; i++ {
		DN = append(DN, pc_1[28+i])
	}
	return DN
}

func (sipher_code *Sipher_code) getCNDN(round int) []int {
	CNDN := sipher_code.getCN(round)
	DN := sipher_code.getDN(round)
	for _, ele := range DN {
		CNDN = append(CNDN, ele)
	}
	return CNDN
}

func (*Sipher_code) replace(CNDN []int) string {

	initial_sipher_code_replace := ""
	for i := 0; i < 56; i++ {
		initial_sipher_code_replace += string(initial_sipher_code[CNDN[i]-1])
	}

	sipher_code_replace := ""
	for i := 0; i < 48; i++ {
		sipher_code_replace += string(initial_sipher_code_replace[pc_2[i]-1])
	}

	return sipher_code_replace
}

func (sipher_code *Sipher_code) GetSipherCodeN(round int) string {
	CNDN := sipher_code.getCNDN(round)
	return sipher_code.replace(CNDN)
}

func getDisplacements(round int) int {
	displacement := 0
	for i := 0; i < round; i++ {
		displacement += displacements[i]
	}
	return displacement
}
