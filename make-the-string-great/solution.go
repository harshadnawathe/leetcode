package makethestringgreat

func toLower(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b
	}
	return 'a' + b - 'A'
}

func isGood(b1, b2 byte) bool {
	if b1 != b2 && toLower(b1) == toLower(b2) {
		return false
	}
	return true
}

// time: O(n)
// space: O(n)
func makeGood(s string) string {
	stk := []byte{}

	for i := 0; i < len(s); i++ {
		b := s[i]
		stk = append(stk, b)

		if len(stk) >= 2 && !isGood(b, stk[len(stk)-2]) {
			stk = stk[:len(stk)-2]
		}
	}

	return string(stk)
}

// time: O(nlogn)
// space: O(n)
func makeGood1(s string) string {
	if len(s) < 2 {
		return s
	}

	mid := len(s) / 2
	s1 := makeGood1(s[:mid])
	s2 := makeGood1(s[mid:])

	return mergeToMakeGood(s1, s2)
}

func mergeToMakeGood(s1, s2 string) string {
	i1 := len(s1) - 1
	i2 := 0

	for i1 >= 0 && i2 < len(s2) && !isGood(s1[i1], s2[i2]) {
		i1--
		i2++
	}

	return s1[:i1+1] + s2[i2:]
}
