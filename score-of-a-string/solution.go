package scoreofastring

func scoreOfString(s string) int {
	score := 0

	for i := range len(s) - 1 {
		score += diff(s[i], s[i+1])
	}

	return score
}

func diff(b1, b2 byte) int {
	if d := int(b1) - int(b2); d < 0 {
		return -d
	} else {
		return d
	}
}
