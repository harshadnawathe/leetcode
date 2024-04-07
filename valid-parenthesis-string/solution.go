package validparenthesisstring

// time: O(n), space: O(1)
func checkValidString(s string) bool {
	numOpened, numWildChar := 0, 0

	for _, char := range s {
		switch char {
		case '(':
			numOpened++
		case '*':
			if numOpened > 0 {
				numOpened--
				numWildChar += 2
			} else {
				numWildChar++
			}
		case ')':
			if numOpened == 0 {
				if numWildChar == 0 {
					return false
				} else {
					numWildChar--
				}
			} else {
				numOpened--
			}
		}
	}

	return numOpened == 0
}
