package sortcolors

const (
	Red = iota
	White
	Blue
)

func sortColors(nums []int) {
	r, w, b := 0, 0, len(nums)-1

	for w <= b {
		switch nums[w] {
		case Red:
			nums[w], nums[r] = nums[r], nums[w]
			w += 1
			r += 1
		case White:
			w += 1
		case Blue:
			nums[w], nums[b] = nums[b], nums[w]
			b -= 1
		}
	}
}
