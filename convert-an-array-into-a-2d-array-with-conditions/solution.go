package convertanarrayintoa2darraywithconditions

import "sort"

// time: O(N) | space: O(N)
func findMatrix(nums []int) [][]int {
	count := make(map[int]int)
	result := [][]int{}

	for _, num := range nums {
		cnt := count[num]
		if len(result) == cnt {
			result = append(result, []int{num})
		} else {
			result[cnt] = append(result[cnt], num)
		}
		count[num]++
	}
	return result
}

// time: O(NlogN) | space: O(1)
func findMatrix1(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{{nums[0]}}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			result[0] = append(result[0], nums[i])
		} else {
			x := 1
			for ; x < len(result); x++ {
				if nums[i] != result[x][len(result[x])-1] {
					break
				}
			}
			if x == len(result) {
				result = append(result, []int{nums[i]})
			} else {
				result[x] = append(result[x], nums[i])
			}
		}
	}

	return result
}
