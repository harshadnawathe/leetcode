package plusone

func plusOne(digits []int) []int {
	carry := true
	for n := len(digits) - 1; n >= 0 && carry; n-- {
		digits[n], carry = inc(digits[n])
	}

	if carry {
		return append([]int{1}, digits...)
	}

	return digits
}

func inc(d int) (int, bool) {
	d++
	if d > 9 {
		return 0, true
	}
	return d, false
}

func plusOne2(digits []int) []int {
	result := []int{0}
	result = append(result, digits...)

	result[len(result)-1] += 1

	for i := len(result) - 1; i > 0 && result[i] > 9; i-- {
		result[i] -= 10
		result[i-1] += 1
	}

	if result[0] != 0 {
		return result
	}

	return result[1:]
}

func plusOne1(digits []int) []int {
	digits[len(digits)-1] += 1

	for i := len(digits) - 1; i > 0 && digits[i] > 9; i-- {
		digits[i] -= 10
		digits[i-1] += 1
	}

	if digits[0] > 9 {
		return append([]int{1, digits[0] - 10}, digits[1:]...)
	}

	return digits
}
