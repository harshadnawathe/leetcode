package optimalpartitionofstring

// Greedy
func partitionString(s string) int {
	if len(s) == 0 {
		return 0
	}
	lastSeen := make(map[byte]int)
	count := 1
	for i, substringStart := 0, 0; i < len(s); i++ {
		if x, ok := lastSeen[s[i]]; ok && x >= substringStart {
			count++
			substringStart = i
		}
		lastSeen[s[i]] = i
	}
	return count
}

func minOf(lhs, rhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}

type bitmap uint32

func mask(x int) uint32 {
	return uint32(1) << uint32(x)
}

func set(b bitmap, x int) bitmap {
	n := uint32(b)
	m := mask(x)
	return bitmap(n | m)
}

func isSet(b bitmap, x int) bool {
	n := uint32(b)
	m := mask(x)
	return n&m != 0
}

func findIf(n int, f func(int) bool) int {
	for i := 0; i < n; i++ {
		if f(i) {
			return i
		}
	}
	return n
}

func isRepeatedIn(s string) (int, func(int) bool) {
	var m bitmap
	return len(s), func(i int) bool {
		k := int(s[i] - 'a')
		if isSet(m, k) {
			return true
		}
		m = set(m, k)
		return false
	}
}

// Recursive with memoization
func partitionString2(s string) int {
	cache := make(map[int]int)

	var minPartition func(int) int
	minPartition = func(i int) int {
		if cached, found := cache[i]; found {
			return cached
		}

		if i == len(s) {
			cache[i] = 0
			return cache[i]
		}

		p := findIf(isRepeatedIn(s[i:]))

		cache[i] = 1 + minOf(minPartition(i+1), minPartition(p+i))
		return cache[i]
	}

	return minPartition(0)
}

// Recursive
func partitionString1(s string) int {
	if len(s) == 0 {
		return 0
	}

	p := findIf(isRepeatedIn(s))

	return 1 + minOf(partitionString1(s[1:]), partitionString1(s[p:]))
}
