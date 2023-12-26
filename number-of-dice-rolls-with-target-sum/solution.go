package numberofdicerollswithtargetsum

const mod = 1000000007

func numRollsToTarget(n int, k int, target int) int {
	cache := make(map[[2]int]int)

	var f func(int, int) int
	f = func(n, target int) int {
		if target < 0 {
			return 0
		}

		if n == 0 {
			if target == 0 {
				return 1
			} else {
				return 0
			}
		}
		key := [2]int{n, target}
		if cached, ok := cache[key]; ok {
			return cached
		}

		count := 0
		for i := 1; i <= k; i++ {
			count += f(n-1, target-i)
			count %= mod
		}
		cache[key] = count
		return count
	}

	return f(n, target)
}
