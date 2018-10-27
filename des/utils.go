package des

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

func getRow(num1, num2 byte) int {
	switch string(num1) + string(num2) {
	case "00":
		return 0
	case "01":
		return 1
	case "10":
		return 2
	case "11":
		return 3
	}
	return -1
}

func getColumn(column string) int {
	switch column {
	case "0000":
		return 0
	case "0001":
		return 1
	case "0010":
		return 2
	case "0011":
		return 3
	case "0100":
		return 4
	case "0101":
		return 5
	case "0110":
		return 6
	case "0111":
		return 7
	case "1000":
		return 8
	case "1001":
		return 9
	case "1010":
		return 10
	case "1011":
		return 11
	case "1100":
		return 12
	case "1101":
		return 13
	case "1110":
		return 14
	case "1111":
		return 15
	}
	return -1
}

func reverse(keys []string) []string {
	keys_reverse := make([]string, 0)
	for i := len(keys) - 1; i >= 0; i-- {
		keys_reverse = append(keys_reverse, keys[i])
	}
	return keys_reverse
}
