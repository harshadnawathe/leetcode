package maxdotproductoftwosubsequences

import "math"

func maxDotProduct(nums1 []int, nums2 []int) int {
	min1, max1 := minmax(nums1)
	min2, max2 := minmax(nums2)

	if max1 < 0 && min2 > 0 {
		return max1 * min2
	}

	if max2 < 0 && min1 > 0 {
		return max2 * min1
	}

	cache := make(map[[2]int]int)
	var mdp func(int, int) int
	mdp = func(i, j int) int {
		if i == len(nums1) || j == len(nums2) {
			return 0
		}
		key := [2]int{i, j}
		if cached, ok := cache[key]; ok {
			return cached
		}

		result := maxOf3(
			nums1[i]*nums2[j]+mdp(i+1, j+1),
			mdp(i, j+1),
			mdp(i+1, j),
		)
		cache[key] = result
		return result
	}

	return mdp(0, 0)
}

func minmax(nums []int) (int, int) {
	var min, max int = math.MaxInt32, math.MinInt32
	for _, num := range nums {
		min = minOf(num, min)
		max = maxOf(num, max)
	}
	return min, max
}

func maxOf3(a, b, c int) int {
	return maxOf(a, maxOf(b, c))
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}

func minOf(lhs, rhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}
