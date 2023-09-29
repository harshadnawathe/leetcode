package monotonicarray

func isMonotonic(nums []int) bool {
	s := zero
	for i := 1; i < len(nums); i++ {
		curr := slopeOf(nums[i-1], nums[i])
		if curr == zero {
			continue
		}
		if s == zero {
			s = curr
			continue
		}
		if curr != s {
			return false
		}
	}
	return true
}

func slopeOf(n1, n2 int) slope {
	if n1 < n2 {
		return up
	}
	if n1 > n2 {
		return down
	}
	return zero
}

type slope int

const (
	zero slope = iota
	up
	down
)
