package minimumnumberofremovalstomakemountainarray

import (
	"cmp"
	"sort"
)

func minimumMountainRemovals(nums []int) int {
	lenLIS := lenLIS(nums)
	lenLDS := lenLDS(nums)

	maxMountainLen := 0

	for i := range nums {
		if lenLIS[i] > 1 && lenLDS[i] > 1 {
			maxMountainLen = max(maxMountainLen, lenLIS[i]+lenLDS[i]-1)
		}
	}

	return len(nums) - maxMountainLen
}

func lenLIS[T cmp.Ordered](a []T) []int {
	lenLIS := make([]int, len(a))

	var lis []T

	for i, elem := range a {
		x := sort.Search(len(lis), func(i int) bool {
			return lis[i] >= elem
		})

		if x < len(lis) {
			lis[x] = elem
		} else {
			lis = append(lis, elem)
		}

		lenLIS[i] = len(lis)
	}

	return lenLIS
}

func lenLDS[T cmp.Ordered](nums []T) []int {
	return reverse(lenLIS(reverse(nums)))
}

func reverse[T any](ts []T) []T {
	begin, end := 0, len(ts)

	for begin < end {
		end--
		ts[begin], ts[end] = ts[end], ts[begin]
		begin++
	}

	return ts
}
