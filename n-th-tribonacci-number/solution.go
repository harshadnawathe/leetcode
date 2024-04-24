package nthtribonaccinumber

func tribonacci(n int) int {
	n0 := 0
	n1 := 1
	n2 := 1

	if n == 0 {
		return n0
	}

	if n == 1 {
		return n1
	}

	if n == 2 {
		return n2
	}

	for i := 3; i <= n; i++ {
		n0, n1, n2 = n1, n2, n0+n1+n2
	}

	return n2

}
