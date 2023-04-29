package findtheduplicatenumber

import (
	"sort"
)

// No modification - no extra memory
// Floyds Algorithm to detect cycle
// Time: O(n)
// Space: O(1)
func findDuplicate(nums []int) int {
	return connectionPointNonTerminating(
		nums[0],
		func(n int) int { return nums[n] },
	)
}

// Algorithm from Elements of Programming : Chapter 2
func collisionPointNonTerminating(x int, f func(int) int) int {
	slow, fast := x, f(x)
	for slow != fast {
		slow = f(slow)
		fast = f(fast)
		fast = f(fast)
	}
	return fast
}

func convergentPointNonTerminating(x0, x1 int, f func(int) int) int {
	for x0 != x1 {
		x0 = f(x0)
		x1 = f(x1)
	}
	return x0
}

func connectionPointNonTerminating(x int, f func(int) int) int {
	return convergentPointNonTerminating(
		x,
		f(collisionPointNonTerminating(x, f)),
		f,
	)
}

// No modification - binary search
// Time: O(nlog(n))
// Space: O(1)
func findDuplicate4(nums []int) int {
	return sort.Search(len(nums), func(i int) bool {
		n := i + 1
		cnt := 0
		for _, num := range nums {
			if num <= n {
				cnt++
			}
		}
		return cnt > n
	}) + 1
}

// Modification - use input as Hashmap
// Time: O(n)
// Space: O(1)
func findDuplicate3(nums []int) int {
	for i := 0; i < len(nums); i++ {
		idx := abs(nums[i]) - 1
		if nums[idx] < 0 {
			return idx + 1
		}
		nums[idx] *= -1
	}
	panic("nums has no duplicate")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// No modification - Hashset - extra memory
// Time: O(n)
// Space: O(n)
func findDuplicate2(nums []int) int {
	set := make(map[int]struct{})

	for _, num := range nums {
		if _, found := set[num]; found {
			return num
		}
		set[num] = struct{}{}
	}
	panic("nums has no duplicate")
}

// Modify array - use sort
// Time: O(nlog(n))
// Space: O(1)
func findDuplicate1(nums []int) int {
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return nums[i-1]
		}
	}

	panic("nums has no duplicate")
}
