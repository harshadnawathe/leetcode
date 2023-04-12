package combinationsum

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	var f func(int, int, []int)
	f = func(target, pos int, combination []int) {
		if target == 0 {
			cp := make([]int, len(combination))
			copy(cp, combination)
			result = append(result, cp)
			return
		}

		for i := pos; i < len(candidates); i++ {
			if target < candidates[i] {
				continue
			}
			combination = append(combination, candidates[i])
			f(target-candidates[i], i, combination)
			combination = combination[:len(combination)-1]
		}
	}

	f(target, 0, []int{})

	return result
}
