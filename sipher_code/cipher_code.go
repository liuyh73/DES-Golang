package sipher_code

import (
    "fmt"
)

var initial_sipher_code = [8][8]int{
    {58, 50, 42, 34, 26, 18, 10, 2},
    {60, 52, 44, 36, 28, 20, 12, 4},
    {62, 54, 46, 38, 30, 22, 14, 6},
    {64, 56, 48, 40, 32, 24, 16, 8},
    {57, 49, 41, 33, 25, 17, 9, 1},
    {59, 51, 43, 35, 27, 19, 11, 3},
    {61, 53, 45, 37, 29, 21, 13, 5},
    {63, 55, 47, 39, 31, 23, 15, 7},
}

var PC_1 = [56]int{
    57, 49, 41, 33, 25, 17, 9,
    1, 58, 50, 42, 34, 26, 18,
    10, 2, 59, 51, 43, 35, 27,
    19, 11, 3, 60, 52, 44, 36,
    63, 55, 47, 39, 31, 23, 15,
    7, 62, 54, 46, 38, 30, 22,
    14, 6, 61, 53, 45, 37, 29,
    21, 13, 5, 28, 20, 12, 4,
}

var PC_2 = [48]int{
    14, 17, 11, 24, 1, 5,
    3, 28, 15, 6, 21, 10,
    23, 19, 12, 4, 26, 8,
    16, 7, 27, 20, 13, 2,
    41, 52, 31, 37, 47, 55,
    30, 40, 51, 45, 33, 48,
    44, 49, 39, 56, 34, 53,
    46, 42, 50, 36, 29, 32
}

var displacements = [16]int{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

type Sipher_code struct {
}

func (*Sipher_code) GetInitialSipherCode() [8][8]int {
    return initial_sipher_code
}

func (*Sipher_code) getCN(round int) [28]int {
    displacement := getDisplacements(round)
    CN := PC_0_sipher_code[displacement:28]
    for i := 0; i < displacement; i++ {
        CN = append(CN, PC_0_sipher_code[i])
    }
    return CN
}

func (*Sipher_code) getDN(round int) [28]int {
    displacement := getDisplacements(round)
    DN := PC_0_sipher_code[28+displacement : 56]
    for i := 0; i < displacement; i++ {
        DN = append(DN, PC_0_sipher_code[28+i])
    }
    return DN
}

func (sipher_code *Sipher_code) GetPCN(round int) [56]int {
    PCN := sipher_code.getCN(round)
    DN := sipher_code.getDN(round)
    for _, ele := range DN {
        PCN = append(PCN, ele)
    }
    return PCN
}

func ()

func getDisplacements(round int) {
    displacement := 0
    for i := 0; i < round; i++ {
        displacement += displacements[i]
    }
    return displacement
}
