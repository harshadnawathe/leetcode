package dividetwointegers

const (
	uintMin = uint32(0)
	uintMax = ^uintMin
	intMax  = int(uintMax >> 1)
	intMin  = -intMax - 1
)

func divide(dividend int, divisor int) int {
	isNegative := (dividend < 0) != (divisor < 0)

	dd, d := abs(dividend), abs(divisor)
	result := 0
	for d <= dd {
		var q uint
		for q = 0; (d << (q + 1)) <= dd; q++ {
		}

		result += 1 << q
		dd -= d << q
	}

	if result > intMax && !isNegative {
		return intMax
	}

	if result < intMin && isNegative {
		return intMin
	}
	if isNegative {
		return -result
	}
	return result
}

func abs(x int) uint {
	if x < 0 {
		return uint(-x)
	}
	return uint(x)
}
