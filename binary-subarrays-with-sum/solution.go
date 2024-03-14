package binarysubarrayswithsum

// time: O(n), space: O(1)
func numSubarraysWithSum(nums []int, goal int) int {
	return slidingWindowAtMost(nums, goal) - slidingWindowAtMost(nums, goal-1)
}

func slidingWindowAtMost(nums []int, goal int) int {
	total := 0

	for start, end, sum := 0, 0, 0; end < len(nums); end++ {

		sum += nums[end]

		for start <= end && sum > goal {
			sum -= nums[start]
			start++
		}

		total += end - start + 1
	}

	return total
}

// // time: O(n), space: O(n)
// func numSubarraysWithSum(nums []int, goal int) int {
// 	sumFreq := make(map[int]int)
// 	total := 0
// 	sumSoFar := 0
// 	for _, num := range nums {
// 		sumSoFar += num
//
// 		if sumSoFar == goal {
// 			total++
// 		}
//
// 		if cnt, ok := sumFreq[sumSoFar - goal]; ok {
//       total += cnt
//     }
//
// 		sumFreq[sumSoFar]++
// 	}
//
// 	return total
// }
//
