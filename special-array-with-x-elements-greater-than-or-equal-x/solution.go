package specialarraywithxelementsgreaterthanorequalx

// time: O(N)
// space: O(N)
func specialArray(nums []int) int {
	count := make([]int, len(nums)+1)

	for _, num := range nums {
		count[min(len(nums), num)]++
	}

	sumGteX := 0

	for x := len(nums); x >= 1; x-- {
		sumGteX += count[x]

		if sumGteX == x {
			return x
		}

	}

	return -1
}
