package countelementswithmaximumfrequency

func maxFrequencyElements(nums []int) int {
	frequency := make(map[int]int)
	max, total := 0, 0

	for _, num := range nums {
		frequency[num] += 1

		f := frequency[num]

		if f > max {
			max = f
			total = f
		} else if f == max {
			total += f
		}
	}

	return total
}
