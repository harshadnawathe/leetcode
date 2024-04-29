package minimumnumberofoperationstomakearrayxorequaltok

func minOperations(nums []int, k int) int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	opCount := 0
	for xor > 0 || k > 0 {
		if xor&0x1 != k&0x1 {
			opCount++
		}
		xor >>= 1
		k >>= 1
	}

	return opCount
}
