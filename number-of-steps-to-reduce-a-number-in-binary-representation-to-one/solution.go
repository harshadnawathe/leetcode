package numberofstepstoreduceanumberinbinaryrepresentationtoone

func numSteps(s string) int {
	bs := []byte(s)

	numSteps := 0

	lsb := len(bs) - 1

	for lsb > 0 {
		if bs[lsb] == '0' {
			lsb--

			numSteps++
		} else {

			p := lsb
			for ; p > 0; p-- {
				if bs[p] == '1' {
					bs[p] = '0'
				} else {
					bs[p] = '1'
					break
				}
			}

			if p != 0 {
				lsb--
			}

			numSteps += 2
		}
	}

	return numSteps
}
