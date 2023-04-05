package removeduplicatesfromsortedarray

func removeDuplicates1(nums []int) int {
	insert := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[insert] = nums[i]
			insert++
		}
	}
	return insert
}

func removeDuplicates(nums []int) int {
	numDup := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			numDup++
		} else {
			nums[i-numDup] = nums[i]
		}
	}
	return len(nums) - numDup
}
