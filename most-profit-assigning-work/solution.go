package mostprofitassigningwork

import (
	"sort"
)

// time: O(nlog(n) + mlog(m))
// space: O(n)
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	sort.Ints(worker)

	jobs := make([]int, len(difficulty))
	for i := range jobs {
		jobs[i] = i
	}

	sort.Slice(jobs, func(i, j int) bool {
		return difficulty[jobs[i]] < difficulty[jobs[j]]
	})

	maxProfit, i, best := 0, 0, 0

	for _, w := range worker {
		for i < len(jobs) && w >= difficulty[jobs[i]] {
			best = max(best, profit[jobs[i]])
			i++
		}
		maxProfit += best
	}

	return maxProfit
}
