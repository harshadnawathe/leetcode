package addstrings

import "unsafe"

func addStrings(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	n := maxOf(n1, n2) + 1
	sum := make([]byte, n)
	sum[0] = '0'

	for n1 > 0 && n2 > 0 {
		n--
		n1--
		n2--
		sum[n] = add(num1[n1], num2[n2])
	}

	for n1 > 0 {
		n--
		n1--
		sum[n] = num1[n1]
	}

	for n2 > 0 {
		n--
		n2--
		sum[n] = num2[n2]
	}

	for i := len(sum) - 1; i > 0; i-- {
		if sum[i] > '9' {
			sum[i] = sub10(sum[i])
			sum[i-1] += 1
		}
	}

	if sum[0] == '0' {
		sum = sum[1:]
	}

	return *(*string)(unsafe.Pointer(&sum))
}

func addStrings1(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	sum := make([]byte, 0, maxOf(n1, n2)+1)

	for n1 > 0 && n2 > 0 {
		n1--
		n2--
		sum = append(sum, add(num1[n1], num2[n2]))
	}

	for n1 > 0 {
		n1--
		sum = append(sum, num1[n1])
	}

	for n2 > 0 {
		n2--
		sum = append(sum, num2[n2])
	}

	sum = append(sum, '0')
	for i := 0; i < len(sum)-1; i++ {
		if sum[i] > '9' {
			sum[i] = sub10(sum[i])
			sum[i+1]++
		}
	}

	reverse(sum)

	if sum[0] == '0' {
		sum = sum[1:]
	}

	return *(*string)(unsafe.Pointer(&sum))
}

func add(b1, b2 byte) byte {
	return b1 + b2 - '0'
}

func sub10(b byte) byte {
	return b - ':' + '0'
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}

func reverse(bs []byte) {
	first, last := 0, len(bs)
	for first < last {
		last--
		bs[first], bs[last] = bs[last], bs[first]
		first++
	}
}
