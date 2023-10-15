package numberofwaystostayinthesameplaceaftersomesteps

const mod = 1000000007

func numWays(steps int, arrLen int) int {
	cache := make(map[[2]int]int)

	var f func(int, int) int
	f = func(pos, step int) int {
		if step == steps {
			if pos == 0 {
				return 1
			}
			return 0
		}

		key := [2]int{pos, step}
		if cached, ok := cache[key]; ok {
			return cached
		}

		cnt := 0
		for _, next := range nextPos(pos, arrLen) {
			cnt += f(next, step+1)
			cnt %= mod
		}
		cache[key] = cnt
		return cnt
	}

	return f(0, 0)
}

func nextPos(pos, arrLen int) []int {
	next := make([]int, 0, 3)
	next = append(next, pos)
	if pos > 0 {
		next = append(next, pos-1)
	}
	if pos < arrLen-1 {
		next = append(next, pos+1)
	}
	return next
}
