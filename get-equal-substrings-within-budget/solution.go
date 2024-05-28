package getequalsubstringswithinbudget

func equalSubstring(s string, t string, maxCost int) int {
	left := 0
	maxLen := 0

	for currCost, right := 0, 0; right < len(s); right++ {
		currCost += convCost(s[right], t[right])

		for currCost > maxCost {
			currCost -= convCost(s[left], t[left])
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

func convCost(b1, b2 byte) int {
	if b1 < b2 {
		return int(b2 - b1)
	}
	return int(b1 - b2)
}
