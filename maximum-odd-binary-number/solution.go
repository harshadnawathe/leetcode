package maximumoddbinarynumber

// time: O(n)
// space: O(n)
func maximumOddBinaryNumber(s string) string {
	bs := []byte(s)
	pp := partition(bs)

	n := len(bs) - 1
	bs[pp-1], bs[n] = bs[n], bs[pp-1]

	return string(bs)
}

func partition(bs []byte) int {
	left, right := 0, len(bs)
	for left < right {
		if bs[left] == '1' {
			left++
			continue
		}
		right--
		for right > left && bs[right] == '0' {
			right--
		}
		if left < right {
			bs[left], bs[right] = bs[right], bs[left]
			left++
		}
	}
	return left
}

// // time: O(n)
// // space: O(n)
// func maximumOddBinaryNumber(s string) string {
//    bs := make([]byte, len(s))
//    p1, p0 := 0, len(s)-1
//
//    for i := 0; i < len(s); i++ {
//      if s[i] == '1' {
//        bs[p1] = '1'
//        p1++
//      } else {
//        bs[p0] = '0'
//        p0--
//      }
//    }
//
//    bs[p1-1] = '0'
//    bs[len(s)-1] = '1'
//
//    return string(bs)
// }
