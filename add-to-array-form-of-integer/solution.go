package addtoarrayformofinteger

func addToArrayForm(num []int, k int) []int {
	result := []int{}

	n := len(num)
	for n > 0 && k > 0 {
		n--
		result = append(result, num[n]+k%10)
		k /= 10
	}

	for n > 0 {
		n--
		result = append(result, num[n])
	}

	for k > 0 {
		result = append(result, k%10)
		k /= 10
	}

	result = append(result, 0)

	for i := 0; i < len(result)-1; i++ {
		if result[i] > 9 {
			result[i] -= 10
			result[i+1] += 1
		}
	}

	reverse(result)

	if result[0] == 0 {
		result = result[1:]
	}

	return result
}

func reverse(a []int) {
	left, right := 0, len(a)
	for left < right {
		right--
		a[left], a[right] = a[right], a[left]
		left++
	}
}
