package removekdigits

func removeKdigits(num string, k int) string {
	stk := []byte{num[0]}

	for i := 1; i < len(num); i++ {

		for k > 0 && len(stk) > 0 && num[i] < stk[len(stk)-1] {
			stk = stk[:len(stk)-1]
			k--
		}

		if len(stk) != 0 || num[i] != '0' {
			stk = append(stk, num[i])
		}
	}

	if n := len(stk) - k; n > 0 {
		stk = stk[:n]
	} else {
		stk = nil
	}

	if len(stk) == 0 {
		return "0"
	}

	return string(stk)
}
