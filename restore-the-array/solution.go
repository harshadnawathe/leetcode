package restorethearray

import "strconv"

const mod int = 1000000007

// tabular dp
func numberOfArrays(s string, k int) int {
	numArrays := make([]int, len(s)+1)
	numArrays[len(s)] = 1

	for begin := len(s) - 1; begin >= 0; begin-- {
		if s[begin] == '0' {
			numArrays[begin] = 0
			continue
		}

		numArrays[begin] = 0
		for end := begin + 1; end <= len(s); end++ {
			n, _ := strconv.Atoi(s[begin:end])
			if n > k {
				break
			}
			numArrays[begin] += numArrays[end]
			numArrays[begin] %= mod

		}
	}

	return numArrays[0]
}

// recursion
func numberOfArrays1(s string, k int) int {
	var numArrays func(int) int
	numArrays = func(begin int) int {
		if begin == len(s) {
			return 1
		}

		if s[begin] == '0' {
			return 0
		}

		count := 0
		for end := begin + 1; end <= len(s); end++ {
			n, _ := strconv.Atoi(s[begin:end])
			if n > k {
				break
			}
			count += numArrays(end)
			count %= mod
		}
		return count
	}
	return numArrays(0)
}
