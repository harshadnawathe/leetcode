package uncrossedlines

// Tabular DP
// Time: O(n1.n2)
// Space: O(n1.n2)
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	maxCount := make([][]int, len(nums1)+1)
	for i := range maxCount {
		maxCount[i] = make([]int, len(nums2)+1)
	}

	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				maxCount[i][j] = 1 + maxCount[i+1][j+1]
			} else {
				maxCount[i][j] = maxOf(
					maxCount[i][j+1],
					maxCount[i+1][j],
				)
			}
		}
	}

	return maxCount[0][0]
}

// Memoization
// Time: O(n1*n2)
// Space: O(n1*n2)
func maxUncrossedLines2(nums1 []int, nums2 []int) int {
	cache := make(map[[2]int]int)
	var f func(int, int) int
	f = func(i, j int) int {
		if i == len(nums1) || j == len(nums2) {
			return 0
		}

		key := [2]int{i, j}
		if cached, found := cache[key]; found {
			return cached
		}

		if nums1[i] == nums2[j] {
			cache[key] = 1 + f(i+1, j+1)
			return cache[key]
		}

		cache[key] = maxOf(
			f(i, j+1),
			f(i+1, j),
		)
		return cache[key]
	}

	return f(0, 0)
}

// Simple recursion
// Time: 2^(n1*n2)
// Space: O(n1*n2) stack
func maxUncrossedLines1(nums1 []int, nums2 []int) int {
	var f func(int, int) int
	f = func(i, j int) int {
		if i == len(nums1) || j == len(nums2) {
			return 0
		}

		if nums1[i] == nums2[j] {
			return 1 + f(i+1, j+1)
		}

		return maxOf(
			f(i, j+1),
			f(i+1, j),
		)
	}
	return f(0, 0)
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}
