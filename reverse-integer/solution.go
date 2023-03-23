package reverseinteger

import "errors"

const (
	uintMin = uint32(0)
	uintMax = ^uintMin
	intMax  = int32(uintMax >> 1)
	intMin  = -intMax - 1

	intMaxBy10      = int(intMax) / 10
	intMaxLastDigit = int(intMax) % 10

	intMinBy10      = int(intMin) / 10
	intMinLastDigit = int(intMin) % 10
)

func reverse(x int) int {
	var reversed int
	for {
		var d int
		x, d = shiftRight(x)

		var err error
		reversed, err = push(reversed, d)
		if err != nil {
			return 0
		}

		if x == 0 {
			break
		}
	}
	return reversed
}

func push(x, d int) (int, error) {

	if x < intMinBy10 || (x == intMinBy10 && d < intMinLastDigit) {
		return 0, errors.New("overflow")
	}
	if x > intMaxBy10 || (x == intMaxBy10 && d > intMaxLastDigit) {
		return 0, errors.New("overflow")
	}

	return x*10 + d, nil
}

func shiftRight(x int) (int, int) {
	return x / 10, x % 10
}
