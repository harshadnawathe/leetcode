package createmaximumnumber

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	max := make([]int, k)
	for i := 0; i <= k; i++ {
		if i <= len(nums1) && (k-i) <= len(nums2) {
			max = maxOf(
				max,
				merge(
					maxNum(nums1, i),
					maxNum(nums2, k-i),
					k,
				),
			)
		}
	}
	return max
}

// maxNum
func maxNum(nums []int, k int) []int {
	if k == len(nums) {
		return nums
	}

	max := stack{}
	drop := len(nums) - k
	for _, num := range nums {
		for drop > 0 && len(max) > 0 && top(max) < num {
			max = pop(max)
			drop--
		}
		max = push(max, num)
	}
	return max[0:k]
}

// merge
func merge(nums1, nums2 []int, k int) []int {
	var merged []int
	for ; k > 0; k-- {
		if len(nums1) > 0 && len(nums2) > 0 {
			if less(nums1, nums2) {
				merged = append(merged, nums2[0])
				nums2 = nums2[1:]
			} else {
				merged = append(merged, nums1[0])
				nums1 = nums1[1:]
			}
		} else if len(nums1) > 0 {
			merged = append(merged, nums1[0])
			nums1 = nums1[1:]
		} else if len(nums2) > 0 {
			merged = append(merged, nums2[0])
			nums2 = nums2[1:]
		}
	}
	return merged
}

// minOf
func minOf(lhs, rhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}

// less
func less(lhs, rhs []int) bool {
	for len(lhs) > 0 && len(rhs) > 0 && lhs[0] == rhs[0] {
		lhs = lhs[1:]
		rhs = rhs[1:]
	}

	return len(lhs) == 0 || (len(rhs) > 0 && lhs[0] < rhs[0])
}

// maxOf
func maxOf(lhs, rhs []int) []int {
	if less(rhs, lhs) {
		return lhs
	}
	return rhs
}

// stack
type stack []int

func top(stk stack) int           { return stk[len(stk)-1] }
func push(stk stack, x int) stack { return append(stk, x) }
func pop(stk stack) stack         { return stk[0 : len(stk)-1] }
