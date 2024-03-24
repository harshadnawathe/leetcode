package takekofeachcharacterfromleftandright

import (
	"sort"
)

// time: O(nlogn)
// soace: O(n)
func takeCharacters(s string, k int) int {
	accFreqLeftAt := []freq{{0, 0, 0}}
	for i := 0; i < len(s); i++ {
		accFreqLeftAt = append(accFreqLeftAt, inc(accFreqLeftAt[i], s[i]))
	}

	leftCharCount := sort.Search(len(accFreqLeftAt), func(i int) bool {
		return hasAtLeastKOfEach(accFreqLeftAt[i], k)
	})

	if leftCharCount == len(accFreqLeftAt) {
		return -1
	}

	minCount := leftCharCount
	accFreqRight := freq{}
	for rightCharCount := 1; rightCharCount <= len(s); rightCharCount++ {
		accFreqRight = inc(accFreqRight, s[len(s)-rightCharCount])

		leftCharCount = sort.Search(leftCharCount, func(i int) bool {
			return hasAtLeastKOfEach(
				add(accFreqLeftAt[i], accFreqRight),
				k,
			)
		})

		if currCount := leftCharCount + rightCharCount; currCount < minCount {
			minCount = currCount
		}
	}

	return minCount
}

type freq [3]int

func inc(c freq, b byte) freq {
	c[b-'a']++
	return c
}

func add(a, b freq) freq {
	return freq{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
	}
}

func hasAtLeastKOfEach(c freq, k int) bool {
	for _, n := range c {
		if n < k {
			return false
		}
	}
	return true
}
