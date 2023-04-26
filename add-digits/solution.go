package adddigits

func addDigits(num int) int {
	if num <= 9 {
		return num
	}

	num %= 9
	if num == 0 {
		return 9
	}

	return num
}
