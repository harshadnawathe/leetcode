package combinations

// Iterative algorithm
func combine(n int, k int) [][]int {
	var combinations [][]int

	combination := make([]int, k)
	i := 0

	for i >= 0 {
		combination[i]++

		if combination[i] > n {
			i--
		} else if i == k-1 {
			copy := append([]int(nil), combination...)
			combinations = append(combinations, copy)
		} else {
			i++
			combination[i] = combination[i-1]
		}
	}

	return combinations
}

// Backtracking algorithm
func combine1(n int, k int) [][]int {
	var combinations [][]int

	var f func(int, []int)
	f = func(i int, combination []int) {
		if len(combination) == k {
			copy := append([]int(nil), combination...)
			combinations = append(combinations, copy)
			return
		}

		for ; i <= n; i++ {
			combination = append(combination, i)
			f(i+1, combination)
			combination = combination[:len(combination)-1]
		}
	}

	f(1, nil)

	return combinations
}
