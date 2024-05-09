package jumpgameii

import "math"

func jump(nums []int) int {
	steps := 0

	far, farthest := 0, 0
	for i, num := range nums {
		if i > far {
			steps++
			far = farthest
		}
		farthest = max(farthest, i+num)
	}

	return steps
}

func jump1(nums []int) int {
	minJumps := make([]int, len(nums))

	for i := len(minJumps) - 1; i >= 0; i-- {
		minJumps[i] = math.MaxInt32
		for j := 1; j <= nums[i]; j++ {
			if i+j < len(minJumps) {
				minJumps[i] = min(minJumps[i], 1+minJumps[i+j])
			}
		}
	}

	return minJumps[0]
}
