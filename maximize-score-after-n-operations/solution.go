package maximizescoreafternoperations

func maxScore(nums []int) int {
	cache := make(map[bitmap]int)
	var f func(bitmap, int) int
	f = func(bm bitmap, n int) int {
		if 2*n == len(nums) {
			return 0
		}
		if cached, found := cache[bm]; found {
			return cached
		}
		max := 0
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				if check(bm, i) || check(bm, j) {
					continue
				}
				next := set(bm, i, j)
				curr := (n + 1) * gcd(nums[i], nums[j])
				remaining := f(next, n+1)

				max = maxOf(max, curr+remaining)
			}
		}
		cache[bm] = max
		return max
	}

	return f(0, 0)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type bitmap uint32

func set(bm bitmap, bits ...int) bitmap {
	result := bm
	for _, bit := range bits {
		result = result | (1 << uint32(bit))
	}
	return result
}

func check(bm bitmap, n int) bool { return (bm & (1 << uint32(n))) != 0 }

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}
