package combinationsumii

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	results := [][]int{}
	var f func(int, int, []int)
	f = func(target, pos int, combination []int) {
		if target < 0 {
			return
		}

		if target == 0 {
			result := make([]int, len(combination))
			copy(result, combination)
			results = append(results, result)
			return
		}

		for i := pos; i < len(candidates); i++ {
			if i > pos && candidates[i] == candidates[i-1] {
				continue
			}
			combination = append(combination, candidates[i])
			f(target-candidates[i], i+1, combination)
			combination = combination[:len(combination)-1]
		}
	}

	sort.Ints(candidates)
	f(target, 0, []int{})

	return results
}
