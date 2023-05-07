package russiandollenvelopes

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	maxLen := 0
	for _, envelope := range envelopes {
		// lower bound
		n := sort.Search(maxLen, func(i int) bool {
			return envelopes[i][1] >= envelope[1]
		})

		if n == maxLen {
			maxLen++
		}

		envelopes[n] = envelope
	}
	return maxLen
}

func maxEnvelopes1(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	var lis []int
	for _, envelope := range envelopes {
		x := sort.Search(len(lis), func(i int) bool {
			return lis[i] >= envelope[1]
		})

		if x == len(lis) {
			lis = append(lis, envelope[1])
		} else {
			lis[x] = envelope[1]
		}
	}

	return len(lis)
}
