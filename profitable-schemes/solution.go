package profitableschemes

const mod int = 1000000007

func minOf(lhs, rhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}

// recursion with memoisation
func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	cache := make(map[[3]int]int)
	var numSchemes func(int, int, int) int
	numSchemes = func(members, total, taskIndex int) int {
		if taskIndex >= len(group) {
			if total >= minProfit {
				return 1
			} else {
				return 0
			}
		}

		key := [3]int{members, total, taskIndex}
		if cached, found := cache[key]; found {
			return cached
		}

		count := numSchemes(members, total, taskIndex+1)
		if members+group[taskIndex] <= n {
			count += numSchemes(
				members+group[taskIndex],
				minOf(minProfit, total+profit[taskIndex]),
				taskIndex+1,
			)
		}
		count %= mod
		cache[key] = count
		return count
	}

	return numSchemes(0, 0, 0)
}

// recursion
func profitableSchemes1(n int, minProfit int, group []int, profit []int) int {
	var numSchemes func(int, int, int) int
	numSchemes = func(members, total, taskIndex int) int {
		if taskIndex >= len(group) {
			if total >= minProfit {
				return 1
			} else {
				return 0
			}
		}

		count := numSchemes(members, total, taskIndex+1)
		if members+group[taskIndex] <= n {
			count += numSchemes(members+group[taskIndex], total+profit[taskIndex], taskIndex+1)
		}
		count %= mod
		return count
	}
	return numSchemes(0, 0, 0)
}
