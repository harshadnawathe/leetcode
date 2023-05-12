package solvingquestionswithbrainpower

// tabular dp
func mostPoints(questions [][]int) int64 {
	maxPoints := make([]int64, len(questions))
	getMaxPoints := func(i int) int64 {
		if i >= len(maxPoints) {
			return 0
		}
		return maxPoints[i]
	}

	for i := len(questions) - 1; i >= 0; i-- {
		pointsIfSolved := int64(questions[i][0]) + getMaxPoints(i+questions[i][1]+1)
		pointsIfSkipped := getMaxPoints(i + 1)

		maxPoints[i] = maxOf(pointsIfSolved, pointsIfSkipped)
	}

	return maxPoints[0]
}

// recursion with memoization
func mostPoints2(questions [][]int) int64 {
	cache := make(map[int]int64)
	var maxPoints func(int) int64
	maxPoints = func(i int) int64 {
		if i >= len(questions) {
			return 0
		}

		if cached, found := cache[i]; found {
			return cached
		}

		pointsIfSolved := int64(questions[i][0]) + maxPoints(i+questions[i][1]+1)
		pointsIfSkipped := maxPoints(i + 1)

		cache[i] = maxOf(pointsIfSolved, pointsIfSkipped)

		return cache[i]
	}
	return maxPoints(0)
}

// recursion
func mostPoints1(questions [][]int) int64 {
	var maxPoints func(int) int64
	maxPoints = func(i int) int64 {
		if i >= len(questions) {
			return 0
		}
		pointsIfSolved := int64(questions[i][0]) + maxPoints(i+questions[i][1]+1)
		pointsIfSkipped := maxPoints(i + 1)
		return maxOf(pointsIfSkipped, pointsIfSolved)
	}
	return maxPoints(0)
}

func maxOf(lhs, rhs int64) int64 {
	if rhs < lhs {
		return lhs
	}
	return rhs
}
