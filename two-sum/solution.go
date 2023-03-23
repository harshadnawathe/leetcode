package twosum

func twoSum(nums []int, target int) []int {
	complement := make(map[int]int)

	for i, num := range nums {
		if j, isPresent := complement[target-num]; isPresent && i != j {
			return []int{i, j}
		}
		complement[num] = i
	}

	return []int{}
}
