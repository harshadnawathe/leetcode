package combinationsum

func combinationSum(candidates []int, target int) [][]int {
	combinations := make([][][]int, target+1)
	for i := range combinations {
		combinations[i] = make([][]int, 0)
	}
	combinations[0] = append(combinations[0], []int{})

	for _, candidate := range candidates {
		for sum := 1; sum < len(combinations); sum++ {
			if sum-candidate < 0 {
				continue
			}

			for _, combination := range combinations[sum-candidate] {
				var newCombination []int
				newCombination = append(newCombination, combination...)
				newCombination = append(newCombination, candidate)

				combinations[sum] = append(combinations[sum], newCombination)
			}
		}
	}

	return combinations[target]
}

func combinationSum1(candidates []int, target int) [][]int {
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
