package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	maxCnt := 0

	for num := range set {
		if _, ok := set[num-1]; ok {
			continue
		}

		cnt := 1

		for {
			if _, ok := set[num+1]; !ok {
				break
			}

			cnt++

			num++
		}

		maxCnt = max(maxCnt, cnt)
	}

	return maxCnt
}
