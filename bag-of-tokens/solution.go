package bagoftokens

import (
	"sort"
)

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)

	faceUp, faceDown := 0, len(tokens)-1
	max, score := 0, 0
	for faceUp <= faceDown {
		if tokens[faceUp] <= power {
			power -= tokens[faceUp]
			faceUp++
			score++
			max = maxOf(max, score)
		} else if score > 0 {
			power += tokens[faceDown]
			faceDown--
			score--
		} else {
			break
		}
	}

	return max
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}
