package nextgreaterelementiii

import (
	"math"
	"sort"
	"strconv"
)

func nextGreaterElement(n int) int {
	bs := []byte(strconv.Itoa(n))

	i := len(bs) - 1
	for ; i > 0 && bs[i] <= bs[i-1]; i-- {
	}

	if i == 0 {
		return -1
	}

	smallest := i
	for j := i + 1; j < len(bs); j++ {
		if bs[j] > bs[i-1] && bs[j] <= bs[smallest] {
			smallest = j
		}
	}

	bs[i-1], bs[smallest] = bs[smallest], bs[i-1]

	sort.Slice(bs[i:], func(x, y int) bool {
		return bs[i+x] < bs[i+y]
	})

	next, err := strconv.Atoi(string(bs))

	if err != nil || next > math.MaxInt32 {
		return -1
	}

	return next
}
