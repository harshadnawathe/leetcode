package numberofstepstoreduceanumberinbinaryrepresentationtoone

func numSteps(s string) int {
	numSteps, carry := 0, 0

	for i := len(s) - 1; i > 0; i-- {
		digit := int(s[i]-'0') + carry

		if digit&0x1 == 1 {
			numSteps += 2
			carry = 1
		} else {
			numSteps++
		}
	}

	return numSteps + carry
}
