package frogjump

type key struct{ i, k int }

func canCross(stones []int) bool {
	stonePos := make(map[int]int)
	for i, stone := range stones {
		stonePos[stone] = i
	}

	cache := make(map[key]bool)

	var f func(int, int) bool
	f = func(i, k int) bool {
		if cached, ok := cache[key{i, k}]; ok {
			return cached
		}

		if i == len(stones)-1 {
			return true
		}

		nextStones := []int{stones[i] + k - 1, stones[i] + k, stones[i] + k + 1}
		for _, next := range nextStones {
			if next == stones[i] {
				continue
			}
			if nextPos, found := stonePos[next]; found {
				if f(nextPos, next-stones[i]) {
					return true
				}
			}
		}

		cache[key{i, k}] = false
		return false
	}

	return f(0, 0)
}
